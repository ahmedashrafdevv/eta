package model

type ItemModel struct {
	Serial        int
	Name          string
	Code          string
	Price         float64
	MinorPerMajor int
}

type ListInvoiceReq struct {
	Finished *bool
	Deleted  *bool
	EmpCode  *int
	DateFrom string `query:"start_date"`
	DateTo   string `query:"end_date"`
}

type LocalInvoiceHead struct {
	Serial         int
	StkTr01Serial  int
	DocNo          int
	DocDate        string
	StcEmpName     string
	EmpCode        int
	DeliveryFee    float64
	DriverName     string
	StoreName      string
	TotalCash      float64
	EmpName        string
	CustomerName   string
	CustomerCode   int
	CustomerSerial int
	Reserved       bool
	Finished       bool
	Deleted        bool
	SalesOrderNo   int
}

type LocalInvoiceDetails struct {
	Serial    int
	BarCode   string
	ItemName  string
	Qnt       float64
	ItemPrice float64
	ItemTotal float64
}

type InsertOrderReq struct {
	DocNo         int
	StoreCode     int
	AccountSerial int
	EmpCode       int
}
type InsertOrderItemReq struct {
	HeadSerial    int
	ItemSerial    int
	ItemName      string
	ItemPrice     float64
	Qnt           float64
	MinorPerMajor float64
	StoreCode     int
}
type UpdateOrderItemReq struct {
	Serial int
	Qnt    float64
}

type GetAccountRequest struct {
	Code *int    `query:"Code"`
	Name *string `query:"Name"`
	Type *int    `query:"Type" validate:"required"`
}

type Account struct {
	Serial      int
	AccountCode int
	AccountName string
}
