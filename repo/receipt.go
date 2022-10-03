package repo

import (
	"eta/model"
	"eta/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

type ReceiptRepo struct {
	db *gorm.DB
}

func NewReceiptRepo(db *gorm.DB) ReceiptRepo {
	return ReceiptRepo{
		db: db,
	}
}

func (ur *ReceiptRepo) ListReceiptsByPosted(posted *bool) (*[]model.EInvoice, error) {
	rows, err := ur.db.Raw("EXEC StkTrEInvoiceHeadList @posted = ? ", posted).Rows()
	utils.CheckErr(&err)
	defer rows.Close()
	if utils.CheckErr(&err) {
		return nil, err
	}
	result, err := scanEInvoiceResult(rows)
	return result, nil
}

func (ur *ReceiptRepo) FindUnPostedReciepts(info *model.CompanyInfo) ([]model.Receipt, error) {
	var resp []model.Receipt
	// var taxRecord model.TaxTotals
	rows, err := ur.db.Raw("EXEC StkTrEInvoiceListReceipts").Rows()
	if utils.CheckErr(&err) {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var rec model.Receipt
		// var taxRecord model.TaxTotals
		// taxRecord.Amount =
		// taxRecord.Amount =
		err := rows.Scan(&rec.Serial, &rec.Header.DateTimeIssued, &rec.Header.ReceiptNumber, &rec.TotalItemsDiscount, &rec.TotalAmount, &rec.TaxTotals.Amount, &rec.Seller.BranchCode, &rec.DocumentType.ReceiptType)
		if utils.CheckErr(&err) {
			return nil, err
		}
		rec.Header.ReceiptNumber = fmt.Sprintf("%s-%d", rec.Seller.BranchCode, rec.Serial)
		rec.Header.Currency = "EGP"
		rec.Buyer.Type = "B"
		rec.DocumentType.ReceiptType = "s"
		rec.DocumentType.TypeVersion = "1.1"
		rec.Seller.Rin = info.EtaRegistrationId
		rec.Seller.CompanyTradeName = info.ComName
		rec.Seller.DeviceSerialNumber = "123"
		rec.TotalItemsDiscount = 0
		rec.ExtraReceiptDiscount = make([]float64, 1)
		rec.PaymentMethod = "C"
		rec.TotalSales = rec.TotalAmount + rec.TaxTotals.Amount
		rec.TaxTotals.TaxType = "T1"
		rec.ExtraReceiptDiscount = append(rec.ExtraReceiptDiscount, 0.0)
		// rec.TaxTotals =
		resp = append(resp, rec)
	}
	if rows.NextResultSet() {
		var headSerial int
		var counter int

		// currentInvoice := invoices[counter]
		for rows.Next() {
			var rec model.ItemData
			var discount float64
			// var head int
			err := rows.Scan(
				&headSerial,
				&rec.Description,
				&rec.ItemType,
				&rec.ItemCode,
				&rec.UnitType,
				&rec.Quantity,
				&rec.UnitPrice,
				&discount,
				&rec.TotalSale,
				&rec.Total,
			)
			// _prepareInvoiceItem(&rec)
			if utils.CheckErr(&err) {
				return nil, err
			}

			if resp[counter].Serial != headSerial {
				counter++
			}
			resp[counter].ItemData = append(resp[counter].ItemData, rec)
		}
	}

	return resp, nil
}
