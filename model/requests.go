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
	Data   []DashboardStoreStats `json:"data"`
	Totals DashboardStoreTotals  `json:"totals"`
}

type DashboardStoreStats struct {
	TotalAmount float64 `json:"totalAmount"`
	TotalTax    float64 `json:"totalTax"`
	NetAmount   float64 `json:"netAmount"`
	StoreCode   int     `json:"storeCode"`
	StoreName   string  `json:"storeName"`
	Orders      int     `json:"orders"`
}

type DashboardStoreTotals struct {
	TotalAmount float64 `json:"totalAmount"`
	TotalTax    float64 `json:"totalTax"`
	NetAmount   float64 `json:"netAmount"`
	Orders      int     `json:"orders"`
}
type ListPosOrdersRequest struct {
	Store    int    `query:"store"`
	Status   int    `query:"status"`
	FromDate string `query:"fromDate"`
	ToDate   string `query:"toDate"`
}
