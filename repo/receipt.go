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
	var serials []int
	// var taxRecord model.TaxTotals
	rows, err := ur.db.Raw("EXEC StkTrEInvoiceListReceipts").Rows()
	if utils.CheckErr(&err) {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var serial int
		var rec model.Receipt
		var taxRecord model.TaxTotals
		// taxRecord.Amount =
		// taxRecord.Amount =
		err := rows.Scan(
			&serial,
			&rec.Header.DateTimeIssued,
			&rec.Header.ReceiptNumber,
			&rec.TotalItemsDiscount,
			&rec.TotalAmount,
			&taxRecord.Amount,
			&rec.Seller.BranchCode,
			&rec.DocumentType.ReceiptType,
			&rec.Seller.BranchAddress.Country,
			&rec.Seller.BranchAddress.Governate,
			&rec.Seller.BranchAddress.RegionCity,
			&rec.Seller.BranchAddress.Street,
			&rec.Seller.BranchAddress.BuildingNumber,
		)
		if utils.CheckErr(&err) {
			return nil, err
		}
		rec.Header.ReceiptNumber = fmt.Sprintf("%s-%d", rec.Seller.BranchCode, serial)
		rec.Header.Currency = "EGP"
		rec.Buyer.Type = "P"
		// rec.DocumentType.ReceiptType = "s"
		rec.DocumentType.TypeVersion = "1.0"
		rec.Seller.Rin = info.EtaRegistrationId
		rec.Seller.CompanyTradeName = info.ComName
		rec.Seller.DeviceSerialNumber = "123"
		rec.TotalItemsDiscount = 0
		rec.PaymentMethod = "C"
		rec.TotalSales = rec.TotalAmount + taxRecord.Amount
		taxRecord.TaxType = "T1"
		rec.Seller.ActivityCode = info.EtaActivityCode
		rec.TaxTotals = make([]model.TaxTotals, 0)
		// rec.ExtraReceiptDiscount = append(rec.ExtraReceiptDiscount, 0.0)
		rec.TaxTotals = append(rec.TaxTotals, taxRecord)
		resp = append(resp, rec)
		serials = append(serials, serial)
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
				&rec.InternalCode,
				&rec.Description,
				&rec.ItemType,
				&rec.ItemCode,
				&rec.UnitType,
				&rec.Quantity,
				&rec.UnitPrice,
				&rec.NetSale,
				&discount,
				&rec.TotalSale,
				&rec.Total,
			)
			if utils.CheckErr(&err) {
				return nil, err
			}
			if serials[counter] != headSerial {
				counter++
			}
			resp[counter].ItemData = append(resp[counter].ItemData, rec)
		}
	}

	return resp, nil
}
