package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	// jwtMiddleware := middleware.JWT(utils.JWTSecret)
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
	invoices := api.Group("/invoices")
	invoices.GET("", h.InvoicesList)
	invoices.POST("/post", h.InvoicePost)

	// receipt routes
	receipt := api.Group("/receipts")
	// receipt.GET("", h.ReceiptsListByPosted)
	receipt.POST("", h.ReceiptPost)

	// global
	api.POST("/upload", h.Upload)
	api.GET("/stores", h.StoresListAll)

}
