package model

type Order struct {
	Serial       int     `json:"serial"`
	DocNo        string  `json:"docNo"`
	DocDate      string  `json:"docDate"`
	Discount     float64 `json:"discount"`
	TotalCash    float64 `json:"totalCash"`
	TotalTax     float64 `json:"totalTax"`
	EtaConverted bool    `json:"etaConverted"`
}
type EInvoice struct {
	Serial              int     `json:"serial"`
	InternalID          string  `json:"internlID"`
	StoreCode           string  `json:"storeCode"`
	TotalDiscountAmount float64 `json:"totalDiscountAmount"`
	TotalAmount         float64 `json:"totalAmount"`
	TotalTax            float64 `json:"totalTax"`
	NetAmount           float64 `json:"netAmount"`

	StkTr01Serial  float64 `json:"stkTr01Serial"`
	DateTimeIssued string  `json:"dateTimeIssued"`
}
