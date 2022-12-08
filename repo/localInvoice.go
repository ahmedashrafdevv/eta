package repo

import (
	"eta/model"
	"eta/utils"

	"github.com/jinzhu/gorm"
)

type LocalInvoiceRepo struct {
	db *gorm.DB
}

func NewLocalInvoiceRepo(db *gorm.DB) LocalInvoiceRepo {
	return LocalInvoiceRepo{
		db: db,
	}
}

func (ir *LocalInvoiceRepo) AccountsList(req *model.GetAccountRequest) (*[]model.Account, error) {
	rows, err := ir.db.Raw("EXEC GetAccount @Code = ?, @Name = ? , @Type = ?", req.Code, req.Name, req.Type).Rows()
	if err != nil {
		return nil, err
	}
	var accounts []model.Account
	defer rows.Close()
	for rows.Next() {
		var account model.Account
		rows.Scan(&account.Serial, &account.AccountCode, &account.AccountName)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)

	}

	return &accounts, nil
}

func (ir *LocalInvoiceRepo) ELocalInvoicesList(req *model.ListInvoiceReq) (*[]model.LocalInvoiceHead, error) {
	rows, err := ir.db.Raw("EXEC StkTrInvoiceHeadList @EmpCode = ? ,   @Finished = ? , @Deleted = ? , @DateFrom = ? , @DateTo = ? ", req.EmpCode, req.Finished, req.Deleted, req.DateFrom, req.DateTo).Rows()
	if err != nil {
		return nil, err
	}
	var orders []model.LocalInvoiceHead
	defer rows.Close()
	for rows.Next() {
		var order model.LocalInvoiceHead
		err = rows.Scan(&order.Serial, &order.DocNo, &order.DocDate, &order.EmpCode, &order.TotalCash, &order.EmpName, &order.CustomerName, &order.CustomerCode, &order.CustomerSerial, &order.Reserved, &order.Finished, &order.Deleted)
		if err != nil {
			return nil, err
		}
		utils.HandleDate(&order.DocDate)
		orders = append(orders, order)
	}

	return &orders, nil
}

func (ir *LocalInvoiceRepo) ELocalInvoicesListItems(invoiceSerial int) (*model.LocalInvoiceResp, error) {
	rows, err := ir.db.Raw("EXEC StkTrInvoiceDetailsList @Serial = ? ", invoiceSerial).Rows()
	if err != nil {
		return nil, err
	}
	var resp model.LocalInvoiceResp
	defer rows.Close()
	for rows.Next() {
		var item model.LocalInvoiceDetails
		err = rows.Scan(&item.Serial, &item.BarCode, &item.ItemName, &item.Qnt, &item.ItemPrice, &item.ItemTotal)
		if err != nil {
			return nil, err
		}
		resp.Items = append(resp.Items, item)
	}

	if rows.NextResultSet() {
		for rows.Next() {
			err := rows.Scan(
				&resp.Totals.SubTotal,
				&resp.Totals.Tax,
				&resp.Totals.Total,
			)
			if utils.CheckErr(&err) {
				return nil, err
			}

		}
	}

	return &resp, nil
}

func (ir *LocalInvoiceRepo) ListItems(store int) (*[]model.ItemModel, error) {
	var resp []model.ItemModel
	rows, err := ir.db.Raw("EXEC GetItems @StoreCode =? ", store).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item model.ItemModel
		rows.Scan(&item.Serial, &item.Name, &item.Code, &item.Price, &item.MinorPerMajor)
		if err != nil {
			return nil, err
		}
		resp = append(resp, item)

	}

	return &resp, nil
}
func (ir *LocalInvoiceRepo) ELocalInvoicesDocNo(transSerial int) (*int, error) {
	var resp int
	err := ir.db.Raw("EXEC StkTrInvoiceDocNo @TrSerial = ?", transSerial).Row().Scan(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
func (ir *LocalInvoiceRepo) ELocalInvoicesOrderInsert(req *model.InsertOrderReq) (*int, error) {
	var result int
	err := ir.db.Raw("EXEC StkTrInvoiceHeadInsert @DocNo = ?, @StoreCode = ? , @EmpCode = ? , @AccountSerial =? ", req.DocNo, req.StoreCode, req.EmpCode, req.AccountSerial).Row().Scan(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ir *LocalInvoiceRepo) ELocalInvoicesOrderItemInsert(req *model.InsertOrderItemReq) (*int, error) {
	var result int
	err := ir.db.Raw("EXEC StkTrInvoiceDetailsInsert @HeadSerial = ?, @ItemSerial = ?, @ItemName = ?, @ItemPrice = ?, @Qnt = ?, @MinorPerMajor = ?, @StoreCode = ?", req.HeadSerial, req.ItemSerial, req.ItemName, req.ItemPrice, req.Qnt, req.MinorPerMajor, req.StoreCode).Row().Scan(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ir *LocalInvoiceRepo) ELocalInvoicesOrderItemUpdate(req *model.UpdateOrderItemReq) (*int, error) {
	var result int
	err := ir.db.Raw("EXEC StkTrInvoiceDetailsUpdate @Serial = ?, @Qnt = ?", req.Serial, req.Qnt).Row().Scan(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ir *LocalInvoiceRepo) ELocalInvoicesOrderClose(serial int) (*int, error) {
	var result int
	err := ir.db.Raw("EXEC StkTrInvoiceOrderClose @Serial = ?", serial).Row().Scan(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ir *LocalInvoiceRepo) ELocalInvoicesOrderItemDelete(serial int) (*int, error) {
	var result int
	err := ir.db.Raw("EXEC StkTrInvoiceDetailsDelete @Serial = ?", serial).Row().Scan(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
