package handler

import (
	"eta/model"
	"eta/repo"
)

type Handler struct {
	userRepo    repo.UserRepo
	orderRepo   repo.OrderRepo
	invoiceRepo repo.InvoiceRepo
	receiptRepo repo.ReceiptRepo
	storeRepo   repo.StoreRepo
	companyInfo *model.CompanyInfo
}

func NewHandler(userRepo repo.UserRepo, orderRepo repo.OrderRepo, invoiceRepo repo.InvoiceRepo, receiptRepo repo.ReceiptRepo, storeRepo repo.StoreRepo, companyInfo *model.CompanyInfo) *Handler {
	return &Handler{
		userRepo:    userRepo,
		orderRepo:   orderRepo,
		invoiceRepo: invoiceRepo,
		receiptRepo: receiptRepo,
		storeRepo:   storeRepo,
		companyInfo: companyInfo,
	}
}
