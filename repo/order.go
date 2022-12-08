package repo

import (
	"database/sql"
	"eta/model"
	"eta/utils"

	"github.com/jinzhu/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return OrderRepo{
		db: db,
	}
}

func (ur *OrderRepo) ListByTransSerialStoreConvertedDate(req *model.ListOrdersRequest) (*[]model.Order, error) {
	rows, err := ur.db.Raw("EXEC StkTr01ListByTransSerialStoreConvertedDate @transSerial = ? , @converted = ? , @storeCode = ?, @fromDate = ? , @toDate = ?", req.TransSerial, req.Status, req.Store, req.FromDate, req.ToDate).Rows()
	if utils.CheckErr(&err) {
		return nil, err
	}
	defer rows.Close()
	result, err := scanOrderResult(rows)
	if utils.CheckErr(&err) {
		return nil, err
	}
	return result, nil
}

func (ur *OrderRepo) DashboardStoreStats(req *model.DashboardStatsRequest) (*model.DashboardStoreStatsResponse, error) {
	var resp model.DashboardStoreStatsResponse
	rows, err := ur.db.Raw("EXEC EtaDashboardStoreStats @start_date = ? , @end_date = ?", req.FromDate, req.ToDate).Rows()
	if utils.CheckErr(&err) {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rec model.DashboardStoreStats
		err := rows.Scan(&rec.TotalAmount, &rec.TotalTax, &rec.StoreName, &rec.StoreCode, &rec.Orders)
		if utils.CheckErr(&err) {
			return nil, err
		}
		rec.NetAmount = rec.TotalAmount - rec.TotalTax
		resp.Totals.TotalAmount += rec.TotalAmount
		resp.Totals.Orders += rec.Orders
		resp.Totals.TotalTax += rec.TotalTax
		resp.Totals.NetAmount += rec.NetAmount
		resp.Data = append(resp.Data, rec)

	}
	return &resp, nil
}

func (ur *OrderRepo) DashboardStats() (*[]model.DashboardStatsResponse, error) {
	var resp []model.DashboardStatsResponse
	rows, err := ur.db.Raw("EXEC EtaDashboardStats").Rows()
	if utils.CheckErr(&err) {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var rec model.DashboardStatsResponse
		err := rows.Scan(&rec.TotalAmount, &rec.TotalTax, &rec.Period)
		if utils.CheckErr(&err) {
			return nil, err
		}
		resp = append(resp, rec)

	}
	return &resp, nil
}

func (ur *OrderRepo) ListPosByStoreConvertedDate(req *model.ListPosOrdersRequest) (*[]model.Order, error) {
	rows, err := ur.db.Raw("EXEC StkTr03ListByStoreConvertedDate  @converted = ? , @storeCode = ? , @fromDate = ? , @toDate = ?", req.Status, req.Store, req.FromDate, req.ToDate).Rows()
	if utils.CheckErr(&err) {
		return nil, err
	}
	defer rows.Close()
	result, err := scanOrderResult(rows)
	if utils.CheckErr(&err) {
		return nil, err
	}
	return result, nil
}

func (ur *OrderRepo) ConvertToEtaInvoice(serial *int64) (*int, error) {
	var resp int
	err := ur.db.Raw("EXEC StkTr01ConvertInvoice @Serial = ? ", serial).Row().Scan(&resp)
	if utils.CheckErr(&err) {
		return nil, err
	}
	return &resp, nil
}

func scanOrderResult(rows *sql.Rows) (*[]model.Order, error) {
	var resp []model.Order
	for rows.Next() {
		var rec model.Order
		err := rows.Scan(&rec.Serial, &rec.DocNo, &rec.DocDate, &rec.Discount, &rec.TotalCash, &rec.TotalTax, &rec.EtaConverted)
		if utils.CheckErr(&err) {
			return nil, err
		}
		resp = append(resp, rec)

	}

	return &resp, nil
}
