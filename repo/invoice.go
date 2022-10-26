package repo

import (
	"database/sql"
	"eta/model"
	"eta/utils"
	"fmt"
	"math"

	"github.com/jinzhu/gorm"
)

type InvoiceRepo struct {
	db *gorm.DB
}

func NewInvoiceRepo(db *gorm.DB) InvoiceRepo {
	return InvoiceRepo{
		db: db,
	}
}

func (ur *InvoiceRepo) ListEInvoices(req *model.ListInvoicessRequest) (*[]model.EInvoice, error) {
	rows, err := ur.db.Raw("EXEC StkTrEInvoiceHeadList @posted = ? , @store = ? , @start_date = ? , @end_date = ? ", req.Posted, req.Store, req.StartDate, req.EndDate).Rows()
	utils.CheckErr(&err)
	defer rows.Close()
	if utils.CheckErr(&err) {
		return nil, err
	}
	result, err := scanEInvoiceResult(rows)
	return result, nil
}

func (ur *InvoiceRepo) RecievedInvoiceInsert(req *model.ReceivedInvoiceInsertReq) (*int, error) {
	var resp int

	err := ur.db.Raw("EXEC EtaRecievedInvoicesHeadInsert  @InternalId = ? , @TotalAmount = ? , @TotalTax = ? , @IssuerName = ?, @IssuerRin = ?, @DateTimeIssued = ? , @DateTimeRecieved = ?   ", req.Invoice.InternalId,
		req.Invoice.TotalAmount,
		req.Invoice.TotalTax,
		req.Invoice.IssuerName,
		req.Invoice.IssuerRin,
		req.Invoice.DateTimeIssued,
		req.Invoice.DateTimeRecieved,
	).Row().Scan(&resp)
	if utils.CheckErr(&err) {
		return nil, err
	}

	for _, v := range req.Items {
		err := ur.db.Raw("EXEC EtaRecievedInvoicesDetailsInsert @InvoiceID = ?  , @ItemName = ? ,@ItemType = ? ,@ItemCode = ? ,@Price = ?  ,@Quantity = ?  ,@TotalAmount = ?  ,@TotalTax = ?  ,@SubTotal = ?    ", resp,
			v.ItemName, v.ItemType, v.ItemCode, v.Price, v.Quantity, v.TotalAmount, v.TotalTax, v.SubTotal,
		).Row().Scan(&resp)
		if utils.CheckErr(&err) {
			return nil, err
		}
	}
	return &resp, nil
}
func (ur *InvoiceRepo) EInvoiceHeadPost(serial *uint64, store *uint64) (*int, error) {
	var resp int
	err := ur.db.Raw("EXEC StkTrEInvoicePosted @serial = ? , @store = ?  ", serial, store).Row().Scan(&resp)

	if utils.CheckErr(&err) {
		return nil, err
	}
	return &resp, nil
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func _removeTax(value *float64) float64 {
	val := *value / 1.14
	return _roundFloat(&val, 5)
}
func _prepareInvoice(info *model.CompanyInfo, invoice *model.Invoice) {
	invoice.Issuer.Id = info.EtaRegistrationId
	internalID := fmt.Sprintf("%s-%d-", invoice.InternalID, invoice.Serial)
	invoice.InternalID = internalID
	invoice.Issuer.Type = info.EtaType
	invoice.TaxpayerActivityCode = info.EtaActivityCode
	invoice.Issuer.Name = info.ComName
	invoice.Receiver.Type = "P"
	invoice.DocumentType = "I"
	invoice.DocumentTypeVersion = "1.0"
	tax := model.TaxTotals{TaxType: "T1", Amount: 0}
	invoice.TaxTotals = append(invoice.TaxTotals, tax)
	invoice.NetAmount = _removeTax(&invoice.TotalAmount)
	invoice.TotalSalesAmount = invoice.NetAmount

}
func _roundFloat(val *float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	roundedValue := math.Round(*val*ratio) / ratio
	return roundedValue
}
func _prepareInvoiceItem(item *model.InvoiceLine) {
	item.NetTotal = item.SalesTotal
	item.UnitValue.CurrencySold = "EGP"
	taxAmount := item.SalesTotal * .14
	taxAmountRounded := _roundFloat(&taxAmount, 5)
	tax := model.TaxableItems{TaxType: "T1", Amount: taxAmountRounded, SubType: " ", Rate: 14}
	item.TaxableItems = append(item.TaxableItems, tax)
	// item.UnitValue.TaxableItems.Ta = "EGP"
}
func (ur *InvoiceRepo) FindInvoiceData(req *model.PostInvoicessRequest, companyInfo *model.CompanyInfo) (*[]model.Invoice, error) {
	var invoices []model.Invoice
	// var invoicesLines []model.InvoiceItem
	rows, err := ur.db.Raw("EXEC StkTrEInvoiceFind @serials = ? , @store = ? ", req.Serilas, req.Store).Rows()
	if utils.CheckErr(&err) {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var rec model.Invoice
		err := rows.Scan(
			&rec.Serial,
			&rec.DateTimeIssued,
			&rec.InternalID,
			&rec.TotalDiscountAmount,
			&rec.TotalAmount,
			&rec.Issuer.Address.BranchId,
			&rec.Issuer.Address.Country,
			&rec.Issuer.Address.Governate,
			&rec.Issuer.Address.RegionCity,
			&rec.Issuer.Address.Street,
			&rec.Issuer.Address.BuildingNumber,
		)
		if utils.CheckErr(&err) {
			return nil, err
		}
		_prepareInvoice(companyInfo, &rec)
		invoices = append(invoices, rec)
	}
	err = ur.db.ScanRows(rows, &invoices)
	if rows.NextResultSet() {
		var headSerial int
		var counter int

		// currentInvoice := invoices[counter]
		for rows.Next() {
			var rec model.InvoiceLine
			// var head int
			err := rows.Scan(
				&headSerial,
				&rec.Description,
				&rec.ItemType,
				&rec.ItemCode,
				&rec.UnitType,
				&rec.Quantity,
				&rec.UnitValue.AmountEGP,
				&rec.ItemsDiscount,
				&rec.SalesTotal,
				&rec.Total,
			)
			_prepareInvoiceItem(&rec)
			if utils.CheckErr(&err) {
				return nil, err
			}

			if invoices[counter].Serial != headSerial {
				counter++
			}
			invoices[counter].InvoiceLines = append(invoices[counter].InvoiceLines, rec)
			invoices[counter].TaxTotals[0].Amount += rec.TaxableItems[0].Amount
			invoices[counter].TaxTotals[0].Amount = _roundFloat(&invoices[counter].TaxTotals[0].Amount, 5)
		}
	}
	return &invoices, nil
}
func scanEInvoiceResult(rows *sql.Rows) (*[]model.EInvoice, error) {
	var resp []model.EInvoice
	for rows.Next() {
		var rec model.EInvoice
		err := rows.Scan(&rec.Serial, &rec.InternalID, &rec.StoreCode, &rec.TotalDiscountAmount, &rec.TotalAmount, &rec.TotalTax, &rec.StkTr01Serial, &rec.DateTimeIssued)
		if utils.CheckErr(&err) {
			return nil, err
		}

		internalID := fmt.Sprintf("%s-%d-", rec.InternalID, rec.Serial)
		rec.InternalID = internalID
		rec.TotalTax = roundFloat(rec.TotalAmount*.14, 5)
		rec.NetAmount = roundFloat((rec.TotalAmount - rec.TotalTax), 5)
		resp = append(resp, rec)
	}
	return &resp, nil
}
