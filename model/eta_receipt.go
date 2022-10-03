package model

type Receipt struct {
	Serial               int          `json:"serial"`
	Header               Header       `json:"header"`
	Seller               Seller       `json:"seller"`
	Buyer                Buyer        `json:"buyer"`
	DocumentType         DocumentType `json:"documentType"`
	ItemData             []ItemData   `json:"itemData"`
	TotalSales           float64      `json:"totalSales"`
	TotalItemsDiscount   float64      `json:"totalItemsDiscount"`
	ExtraReceiptDiscount []float64    `json:"extraReceiptDiscount"`
	NetAmount            float64      `json:"netAmount"`
	FeesAmount           float64      `json:"feesAmount"`
	TotalAmount          float64      `json:"totalAmount"`
	TaxTotals            TaxTotals    `json:"taxTotals"`
	PaymentMethod        string       `json:"paymentMethod"`
}

type DocumentType struct {
	ReceiptType string `json:"receiptType"`
	TypeVersion string `json:"typeVersion"`
}
type Header struct {
	DateTimeIssued string `json:"dateTimeIssued"`
	ReceiptNumber  string `json:"receiptNumber"`
	Uuid           string `json:"uuid"`
	PreviousUUID   string `json:"previousUUID"`
	Currency       string `json:"currency"`
}
type Seller struct {
	Rin              string `json:"rin"`
	CompanyTradeName string `json:"companyTradeName"`
	BranchCode       string `json:"branchCode"`
	// BranchAddress      RecieverAddress `json:"branchAddress"`
	DeviceSerialNumber string `json:"deviceSerialNumber"`
}
type Buyer struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}
type ItemData struct {
	InternalCode string  `json:"internalCode"`
	Description  string  `json:"description"`
	ItemType     string  `json:"itemType"`
	ItemCode     string  `json:"itemCode"`
	UnitType     string  `json:"unitType"`
	Quantity     float64 `json:"quantity"`
	UnitPrice    float64 `json:"unitPrice"`
	NetSale      float64 `json:"netSale"`
	TotalSale    float64 `json:"totalSale"`
	Total        float64 `json:"total"`
}
