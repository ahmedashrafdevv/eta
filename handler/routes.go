package handler

import (
	"eta/router/middleware"
	"eta/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
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
	invoices.GET("/local", h.InvoicesLocalList, jwtMiddleware)
	invoices.GET("/local/items/:id", h.InvoicesLocalListItems, jwtMiddleware)
	invoices.DELETE("/local/items/:id", h.InvoicesLocalDeleteItem, jwtMiddleware)
	invoices.PUT("/local/items", h.InvoicesLocalUpdateItem, jwtMiddleware)
	invoices.PUT("/local/close/:id", h.InvoicesLocalClose, jwtMiddleware)
	invoices.GET("/local/no", h.InvoicesLocalDocNo, jwtMiddleware)
	invoices.POST("/local", h.InvoicesLocalOrderInsert, jwtMiddleware)
	invoices.POST("/local/item", h.InvoicesLocalOrderItemInsert, jwtMiddleware)

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
