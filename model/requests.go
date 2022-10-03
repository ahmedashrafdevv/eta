package model

type ListOrdersRequest struct {
	Store       int    `query:"store"`
	Status      *int   `query:"active"`
	TransSerial int    `query:"transSerial"`
	FromDate    string `query:"fromDate"`
	ToDate      string `query:"toDate"`
}

type ListPosOrdersRequest struct {
	Store    int    `query:"store"`
	Status   int    `query:"status"`
	FromDate string `query:"fromDate"`
	ToDate   string `query:"toDate"`
}
