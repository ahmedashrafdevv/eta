package model

type ListOrdersRequest struct {
	Store       int    `query:"store"`
	Status      *int   `query:"active"`
	TransSerial int    `query:"transSerial"`
	FromDate    string `query:"fromDate"`
	ToDate      string `query:"toDate"`
}

type DashboardStatsRequest struct {
	FromDate *string `query:"start_date"`
	ToDate   *string `query:"end_date"`
}

type DashboardStatsResponse struct {
	TotalAmount float64 `json:"totalAmount"`
	TotalTax    float64 `json:"totalTax"`
	Period      string  `json:"period"`
}

type DashboardStoreStatsResponse struct {
	TotalAmount float64 `json:"totalAmount"`
	TotalTax    float64 `json:"totalTax"`
	StoreCode   int     `json:"storeCode"`
	StoreName   string  `json:"storeName"`
}
type ListPosOrdersRequest struct {
	Store    int    `query:"store"`
	Status   int    `query:"status"`
	FromDate string `query:"fromDate"`
	ToDate   string `query:"toDate"`
}
