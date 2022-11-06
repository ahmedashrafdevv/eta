package model

type EtaLoginRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

type EtaLoginResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}
type EtaRecentDocumentsResponse struct {
	Result []EtaRecentDocumentsItem `json:"result"`
}
type ListReceivedReq struct {
	FromDate *string `query:"start_date"`
	ToDate   *string `query:"end_date"`
	Rin      *string `query:"rin"`
}

type ListReceivedResp struct {
	Id               int
	UUID             string
	InternalId       string
	TotalAmount      float64
	TotalTax         float64
	IssuerName       string
	IssuerRin        string
	DateTimeIssued   string
	DateTimeReceived string
}

type ListReceivedItemsResp struct {
	Id          int
	ItemName    string
	ItemType    string
	ItemCode    string
	ItemPrice   float64
	Quantity    float64
	TotalAmount float64
	TotalTax    float64
	SubTotal    float64
	InvoiceId   int
}
type EtaRecentDocumentsItem struct {
	PublicUrl                     string         `json:"publicUrl"`
	Uuid                          string         `json:"uuid"`
	SubmissionUUID                string         `json:"submissionUUID"`
	LongId                        string         `json:"longId"`
	InternalId                    string         `json:"internalId"`
	TypeName                      string         `json:"typeName"`
	DocumentTypeNamePrimaryLang   string         `json:"documentTypeNamePrimaryLang"`
	DocumentTypeNameSecondaryLang string         `json:"documentTypeNameSecondaryLang"`
	TypeVersionName               string         `json:"typeVersionName"`
	IssuerId                      string         `json:"issuerId"`
	IssuerName                    string         `json:"issuerName"`
	ReceiverId                    string         `json:"receiverId"`
	ReceiverName                  string         `json:"receiverName"`
	DateTimeIssued                string         `json:"dateTimeIssued"`
	DateTimeReceived              string         `json:"dateTimeReceived"`
	TotalSales                    float64        `json:"totalSales"`
	TotalTax                      float64        `json:"totalTax"`
	TotalDiscount                 float64        `json:"totalDiscount"`
	NetAmount                     float64        `json:"netAmount"`
	Total                         float64        `json:"total"`
	MaxPercision                  float64        `json:"maxPercision"`
	InvoiceLineItemCodes          string         `json:"invoiceLineItemCodes"`
	CancelRequestDate             string         `json:"cancelRequestDate"`
	RejectRequestDate             string         `json:"rejectRequestDate"`
	CancelRequestDelayedDate      string         `json:"cancelRequestDelayedDate"`
	RejectRequestDelayedDate      string         `json:"rejectRequestDelayedDate"`
	DeclineCancelRequestDate      string         `json:"declineCancelRequestDate"`
	DeclineRejectRequestDate      string         `json:"declineRejectRequestDate"`
	DocumentStatusReason          string         `json:"documentStatusReason"`
	Status                        string         `json:"status"`
	CreatedByUserId               string         `json:"createdByUserId"`
	FreezeStatus                  EtaFreezStatus `json:"freezeStatus"`
}

type EtaFreezStatus struct {
	Frozen     bool   `json:"frozen"`
	Type       string `json:"type"`
	ActionDate string `json:"actionDate"`
	AuCode     string `json:"auCode"`
	AuName     string `json:"auName"`
}
type EtaSubmitInvoiceResponse struct {
	SubmissionId      string            `json:"submissionId"`
	AcceptedDocuments string            `json:"acceptedDocuments"`
	RejectedDocuments string            `json:"rejectedDocuments"`
	Header            EtaResponseHeader `json:"header"`
}

// "statusCode": "Success",
// "code": "Success",
// "details": [],
// "correlationId": "0HML0CC2QJGUP:00000002",
// "requestTime": "2022-10-12T14:28:15.7268276Z",
// "responseProcessingTimeInTicks": 2775960
type EtaResponseHeader struct {
	StatusCode                    string `json:"statusCode"`
	Code                          string `json:"code"`
	Details                       string `json:"details"`
	CorrelationId                 string `json:"correlationId"`
	RequestTime                   string `json:"requestTime"`
	ResponseProcessingTimeInTicks int    `json:"responseProcessingTimeInTicks"`
}
type EtaResponseRejected struct {
	ReceiptNumber string             `json:"receiptNumber"`
	Uuid          string             `json:"uuid"`
	Error         []EtaResponseError `json:"error"`
}
type EtaResponseError struct {
	Message string                    `json:"message"`
	Target  string                    `json:"target"`
	Details []EtaResponseErrorDetails `json:"details"`
}
type EtaResponseErrorDetails struct {
	Message      string `json:"message"`
	Target       string `json:"target"`
	PropertyPath string `json:"propertyPath"`
}

// "receiptNumber": "7-29288",
//
//	"uuid": "6766f5162ae8ea0444e92e3850c8cb0fdd7b1480744b22a4566fef2c5c792e2b",
//	"error": {
//	    "message": "Validation Error",
//	    "target": "6766f5162ae8ea0444e92e3850c8cb0fdd7b1480744b22a4566fef2c5c792e2b",
//	    "details": [
//	        {
//	            "message": "NoAdditionalPropertiesAllowed",
//	            "target": "referenceUUID",
//	            "propertyPath": "#/header.referenceUUID"
//	        }
//	    ]
//	}
type EtaSubmitInvoiceFailedResponse struct {
	Header            EtaResponseHeader     `json:"header"`
	RejectedDocuments []EtaResponseRejected `json:"rejectedDocuments"`
}

type EtaInvoiceRejectBody struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}
type EtaInvoiceDetailsResp struct {
	SubmissionId                  string               `json:"submissionId"`
	DateTimeRecevied              string               `json:"dateTimeRecevied"`
	ValidationResults             ValidationResult     `json:"validationResults"`
	TransformationStatus          string               `json:"transformationStatus"`
	StatusId                      int                  `json:"statusId"`
	Status                        string               `json:"status"`
	DocumentStatusReason          string               `json:"documentStatusReason"`
	CancelRequestDate             string               `json:"cancelRequestDate"`
	RejectRequestDate             string               `json:"rejectRequestDate"`
	CancelRequestDelayedDate      string               `json:"cancelRequestDelayedDate"`
	RejectRequestDelayedDate      string               `json:"rejectRequestDelayedDate"`
	DeclineCancelRequestDate      string               `json:"declineCancelRequestDate"`
	DeclineRejectRequestDate      string               `json:"declineRejectRequestDate"`
	CanbeCancelledUntil           string               `json:"canbeCancelledUntil"`
	CanbeRejectedUntil            string               `json:"canbeRejectedUntil"`
	SubmissionChannel             int                  `json:"submissionChannel"`
	FreezeStatus                  EtaFreezStatus       `json:"freezeStatus"`
	Uuid                          string               `json:"uuid"`
	PublicUrl                     string               `json:"publicUrl"`
	PurchaseOrderDescription      string               `json:"purchaseOrderDescription"`
	TotalItemsDiscountAmount      float64              `json:"totalItemsDiscountAmount"`
	Delivery                      Delivery             `json:"delivery"`
	Payment                       EtaPayment           `json:"payment"`
	TotalAmount                   float64              `json:"totalAmount"`
	TaxTotals                     []TaxTotals          `json:"taxTotals"`
	NetAmount                     float64              `json:"netAmount"`
	TotalDiscount                 float64              `json:"totalDiscount"`
	TotalSales                    float64              `json:"totalSales"`
	InvoiceLines                  []DetailsInvoiceLine `json:"invoiceLines"`
	References                    []string             `json:"references"`
	SalesOrderDescription         string               `json:"salesOrderDescription"`
	SalesOrderReference           string               `json:"salesOrderReference"`
	ProformaInvoiceNumber         string               `json:"proformaInvoiceNumber"`
	Signatures                    []DetailsSignature   `json:"signatures"`
	PurchaseOrderReference        string               `json:"purchaseOrderReference"`
	InternalID                    string               `json:"internalID"`
	TaxpayerActivityCode          string               `json:"taxpayerActivityCode"`
	DateTimeIssued                string               `json:"dateTimeIssued"`
	DocumentTypeVersion           string               `json:"documentTypeVersion"`
	DocumentType                  string               `json:"documentType"`
	DocumentTypeNamePrimaryLang   string               `json:"documentTypeNamePrimaryLang"`
	DocumentTypeNameSecondaryLang string               `json:"documentTypeNameSecondaryLang"`
	Receiver                      DetailsReceiver      `json:"receiver"`
	Issuer                        DetailsIssuer        `json:"issuer"`
	ExtraDiscountAmount           float64              `json:"extraDiscountAmount"`
	MaxPercision                  int                  `json:"maxPercision"`
	CurrenciesSold                string               `json:"currenciesSold"`
	CurrencySegments              []CurrencySegment    `json:"currencySegments"`
}

type Delivery struct {
	Approach        string  `json:"approach"`
	Packaging       string  `json:"packaging"`
	DateValidity    string  `json:"dateValidity"`
	ExportPort      string  `json:"exportPort"`
	CountryOfOrigin string  `json:"countryOfOrigin"`
	GrossWeight     float64 `json:"grossWeight"`
	NetWeight       float64 `json:"netWeight"`
	Terms           string  `json:"terms"`
}
type CurrencySegment struct {
	Currency                 string      `json:"currency"`
	CurrencyExchangeRate     float64     `json:"currencyExchangeRate"`
	TotalItemsDiscountAmount float64     `json:"totalItemsDiscountAmount"`
	TotalAmount              float64     `json:"totalAmount"`
	TaxTotals                []TaxTotals `json:"taxTotals"`
	NetAmount                float64     `json:"netAmount"`
	TotalDiscount            float64     `json:"totalDiscount"`
	ExtraDiscountAmount      float64     `json:"extraDiscountAmount"`
	TotalSales               float64     `json:"totalSales"`
	TotalTaxableFees         float64     `json:"totalTaxableFees"`
}
type DetailsIssuer struct {
	Type    int             `json:"type"`
	Id      string          `json:"id"`
	Name    string          `json:"name"`
	Address ReceiverAddress `json:"address"`
}
type DetailsReceiver struct {
	Type    int             `json:"type"`
	Id      string          `json:"id"`
	Name    string          `json:"name"`
	Address ReceiverAddress `json:"address"`
}

type EtaReceiverAddress struct {
	BuildingNumber        string `json:"buildingNumber"`
	Room                  string `json:"room"`
	Floor                 string `json:"floor"`
	Street                string `json:"street"`
	Landmark              string `json:"landmark"`
	AdditionalInformation string `json:"additionalInformation"`
	Governate             string `json:"governate"`
	RegionCity            string `json:"regionCity"`
	Country               string `json:"country"`
	BranchId              string `json:"branchID"`
}
type DetailsSignature struct {
	SignatureType string `json:"signatureType"`
	Value         string `json:"value"`
	SignedBy      string `json:"signedBy"`
}

type DetailsInvoiceLine struct {
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
type EtaPayment struct {
	BankName        string `json:"bankName"`
	BankAddress     string `json:"bankAddress"`
	BankAccountNo   string `json:"bankAccountNo"`
	BankAccountIBAN string `json:"bankAccountIBAN"`
	SwiftCode       string `json:"swiftCode"`
	Terms           string `json:"terms"`
}

type ValidationResult struct {
	Status          string           `json:"status"`
	ValidationSteps []ValidationStep `json:"validationSteps"`
}

type ValidationStep struct {
	Status   string `json:"status"`
	Error    string `json:"error"`
	StepName string `json:"stepName"`
	StepId   string `json:"stepId"`
}
