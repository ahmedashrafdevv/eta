package handler

import (
	"eta/model"
	"eta/repo"
)

type Handler struct {
	userRepo         repo.UserRepo
	orderRepo        repo.OrderRepo
	invoiceRepo      repo.InvoiceRepo
	localInvoiceRepo repo.LocalInvoiceRepo
	receiptRepo      repo.ReceiptRepo
	storeRepo        repo.StoreRepo
	companyInfo      *model.CompanyInfo
}

func NewHandler(userRepo repo.UserRepo, orderRepo repo.OrderRepo, invoiceRepo repo.InvoiceRepo, localInvoiceRepo repo.LocalInvoiceRepo, receiptRepo repo.ReceiptRepo, storeRepo repo.StoreRepo, companyInfo *model.CompanyInfo) *Handler {
	return &Handler{
		userRepo:         userRepo,
		orderRepo:        orderRepo,
		invoiceRepo:      invoiceRepo,
		localInvoiceRepo: localInvoiceRepo,
		receiptRepo:      receiptRepo,
		storeRepo:        storeRepo,
		companyInfo:      companyInfo,
	}
}
