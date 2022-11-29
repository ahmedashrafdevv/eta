package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	// jwtMiddleware := middleware.JWT(utils.JWTSecret)
	// auth := v1.Group("/", jwtMiddleware)

	v1.GET("/health", h.CheckHealth)
	api := v1.Group("/api")
	api.POST("/test", h.Test)
	api.GET("/validate", h.ValidateUser)
	//auth routes
	v1.POST("api/login", h.Login)
	// cuurent user routes
	currentUser := api.Group("/me")
	currentUser.GET("", h.Me)

	// orders routes
	orders := api.Group("/orders")
	orders.POST("/convert/:serial", h.OrdersConvertToEta)
	orders.GET("", h.OrdersListByTransSerialStoreConvertedDate)

	// invoices routes
	dashboard := api.Group("/dashboard")
	dashboard.GET("", h.DashboardStats)
	dashboard.GET("/store", h.DashboardStoreStats)

	invoices := api.Group("/invoices")
	invoices.GET("", h.InvoicesList)
	invoices.GET("/recent", h.RecivedInvoicesList)
	invoices.GET("/recent/:id", h.RecievedInvoiceListItems)
	invoices.PUT("/recent/reject/:id", h.RecievedInvoiceReject)
	invoices.POST("/post", h.InvoicePost)
	invoices.POST("/etl", h.InvoicesRecentEtl)

	// Local invoices
	invoices.GET("/local", h.InvoicesLocalList)
	invoices.GET("/local/items/:id", h.InvoicesLocalListItems)
	invoices.DELETE("/local/items/:id", h.InvoicesLocalDeleteItem)
	invoices.PUT("/local/items", h.InvoicesLocalUpdateItem)
	invoices.GET("/local/no", h.InvoicesLocalDocNo)
	invoices.POST("/local", h.InvoicesLocalOrderInsert)
	invoices.POST("/local/item", h.InvoicesLocalOrderItemInsert)

	// receipt routes
	receipt := api.Group("/receipts")
	// receipt.GET("", h.ReceiptsListByPosted)
	receipt.POST("", h.ReceiptPost)

	// global
	api.POST("/upload", h.Upload)
	api.GET("/stores", h.StoresListAll)
	api.GET("/accounts", h.AccountsList)
	api.GET("/items", h.ItemsList)

}
