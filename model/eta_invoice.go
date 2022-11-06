package model

type CompanyInfo struct {
	EtaRegistrationId string
	EtaActivityCode   string
	EtaType           string
	ComName           string
}

type ListInvoicessRequest struct {
	StartDate *string `query:"start_date"`
	EndDate   *string `query:"end_date"`
	Posted    *bool   `query:"posted"`
	Store     *int    `query:"store"`
}

type PostInvoicessRequest struct {
	Serilas string `json:"serials"`
	Store   int    `query:"store"`
}

type Issuer struct {
	Type    string        `json:"type"`
	Id      string        `json:"id"`
	Name    string        `json:"name"`
	Address IssuerAddress `json:"address"`
}

type TaxableItems struct {
	TaxType string  `json:"taxType"`
	Amount  float64 `json:"amount"`
	SubType string  `json:"subType"`
	Rate    float64 `json:"rate"`
}

type ReceiverAddress struct {
	Country        string `json:"country"`
	Governate      string `json:"governate"`
	RegionCity     string `json:"regionCity"`
	Street         string `json:"street"`
	BuildingNumber string `json:"buildingNumber"`
}
type IssuerAddress struct {
	BranchId       string `json:"branchId"`
	Country        string `json:"country"`
	Governate      string `json:"governate"`
	RegionCity     string `json:"regionCity"`
	Street         string `json:"street"`
	BuildingNumber string `json:"buildingNumber"`
}
type Receiver struct {
	Type string `json:"type"`
	// Address ReceiverAddress `json:"address"`
}

type InvoiceHead struct {
	DateTimeIssued      string  `json:"DateTimeIssued"`
	InternalID          string  `json:"internalID"`
	TotalDiscountAmount float64 `json:"totalDiscountAmount"`
	TotalAmount         float64 `json:"totalAmount"`
	AccountSerial       int     `json:"accountSerial"`
	StoreCode           int     `json:"storeCode"`
	EtaStoreCode        string  `json:"etaStoreCode"`
	Country             string  `json:"country"`
	Governate           string  `json:"governate"`
	RegionCity          string  `json:"regionCity"`
	Street              string  `json:"street"`
	BuildingNumber      string  `json:"buildingNumber"`
}

type InvoiceItem struct {
	HeadSerial    int     `json:"HeadSerial"`
	ItemType      string  `json:"itemType"`
	ItemCode      string  `json:"itemCode"`
	UnitType      string  `json:"unitType"`
	Quantity      float64 `json:"quantity"`
	UnitValue     float64 `json:"unitValue"`
	ItemsDiscount float64 `json:"itemsDiscount"`
	SalesTotal    float64 `json:"salesTotal"`
	Total         float64 `json:"total"`
}
type InvoiceLine struct {
	Description      string         `json:"description"`
	ItemType         string         `json:"itemType"`
	ItemCode         string         `json:"itemCode"`
	UnitType         string         `json:"unitType"`
	Quantity         float64        `json:"quantity"`
	SalesTotal       float64        `json:"salesTotal"`
	Total            float64        `json:"total"`
	ValueDifference  float64        `json:"valueDifference"`
	TotalTaxableFees float64        `json:"totalTaxableFees"`
	NetTotal         float64        `json:"netTotal"`
	ItemsDiscount    float64        `json:"itemsDiscount"`
	UnitValue        Value          `json:"unitValue"`
	TaxableItems     []TaxableItems `json:"taxableItems"`
}

type TaxTotals struct {
	TaxType string  `json:"taxType"`
	Amount  float64 `json:"amount"`
}
type Value struct {
	CurrencySold string  `json:"currencySold"`
	AmountEGP    float64 `json:"amountEGP"`
}
type Signature struct {
	SignatureType string `json:"signatureType"`
	Value         string `json:"value"`
}

type Invoice struct {
	Serial                   int
	Issuer                   Issuer        `json:"issuer"`
	Receiver                 Receiver      `json:"receiver"`
	DocumentType             string        `json:"documentType"`
	DocumentTypeVersion      string        `json:"documentTypeVersion"`
	DateTimeIssued           string        `json:"dateTimeIssued"`
	TaxpayerActivityCode     string        `json:"taxpayerActivityCode"`
	InternalID               string        `json:"internalID"`
	InvoiceLines             []InvoiceLine `json:"invoiceLines"`
	TotalDiscountAmount      float64       `json:"totalDiscountAmount"`
	TotalItemsDiscountAmount float64       `json:"totalItemsDiscountAmount"`
	NetAmount                float64       `json:"netAmount"`
	TotalSalesAmount         float64       `json:"totalSalesAmount"`
	ExtraDiscountAmount      float64       `json:"extraDiscountAmount"`
	TotalAmount              float64       `json:"totalAmount"`
	TaxTotals                []TaxTotals   `json:"taxTotals"`
}

type SignedInvoice struct {
	Issuer                   Issuer        `json:"issuer"`
	Receiver                 Receiver      `json:"receiver"`
	DocumentType             string        `json:"documentType"`
	DocumentTypeVersion      string        `json:"documentTypeVersion"`
	DateTimeIssued           string        `json:"dateTimeIssued"`
	TaxpayerActivityCode     string        `json:"taxpayerActivityCode"`
	InternalID               string        `json:"internalID"`
	InvoiceLines             []InvoiceLine `json:"invoiceLines"`
	TotalDiscountAmount      float64       `json:"totalDiscountAmount"`
	TotalItemsDiscountAmount float64       `json:"totalItemsDiscountAmount"`
	NetAmount                float64       `json:"netAmount"`
	TotalSalesAmount         float64       `json:"totalSalesAmount"`
	ExtraDiscountAmount      float64       `json:"extraDiscountAmount"`
	TotalAmount              float64       `json:"totalAmount"`
	TaxTotals                []TaxTotals   `json:"taxTotals"`
	Signatures               []Signature   `json:"signatures"`
}

type InvoiceSubmitRequest struct {
	Documents []SignedInvoice `json:"documents"`
}

type ReceiptSubmitRequest struct {
	Receipts []Receipt `json:"receipts"`
}

type ReceivedInvoiceInsertReq struct {
	Invoice ReceivedInvoice
	Items   []EtaRecentDocumentsItem
}
type ReceivedInvoice struct {
	InternalId       string
	TotalAmount      float64
	TotalTax         float64
	IssuerName       string
	IssuerRin        string
	DateTimeIssued   string
	DateTimeRecieved string
}

// add three names on item name
type ReceivedInvoiceItem struct {
	ItemName    string
	ItemType    string
	ItemCode    string
	Price       float64
	Quantity    float64
	TotalAmount float64
	TotalTax    float64
	SubTotal    float64
}
type InvoiceFindResp struct {
	Serial                   int
	DateTimeIssued           string
	InternalId               string
	TotalSalesAmount         float64
	TotalItemsDiscountAmount float64
	TotalDiscountAmount      float64
	ExtraDiscountAmount      float64
	NetAmount                float64
	TotalAmount              float64
}
