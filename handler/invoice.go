package handler

import (
	"eta/model"
	"eta/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) InvoicesRecentEtl(c echo.Context) error {
	d, err := utils.EtaRecentDocuments()
	type resultStruct struct {
		invoice model.EtaRecentDocumentsItem
		items   []model.DetailsInvoiceLine
	}
	var result []resultStruct
	// var req []model.ReceivedInvoiceInsertReq
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusOK, err.Error())
	}
	for _, v := range d.Result {
		if v.IssuerId != "288271998" {
			existed := h.invoiceRepo.RecievedInvoiceCheckExists(&v.Uuid)
			if !existed {
				v.TotalTax = v.Total - v.TotalSales
				utils.HandleDate(&v.DateTimeIssued)
				utils.HandleDate(&v.DateTimeReceived)

				v.TotalTax = utils.RoundFloat(&v.TotalTax, 5)
				parentId, err := h.invoiceRepo.RecievedInvoiceHeadInsert(v)
				if utils.CheckErr(&err) {
					return c.JSON(http.StatusOK, err.Error())
				}
				invoiceDetails, err := utils.EtaRecentDocumentView(v.Uuid)
				if utils.CheckErr(&err) {
					return c.JSON(http.StatusOK, err.Error())
				}

				fmt.Println("inv")
				result = append(result, resultStruct{invoice: v, items: invoiceDetails.InvoiceLines})
				for _, item := range invoiceDetails.InvoiceLines {
					fmt.Print('1')
					_, err = h.invoiceRepo.RecievedInvoiceDetailsInsert(parentId, item)
					if utils.CheckErr(&err) {
						return c.JSON(http.StatusOK, err.Error())
					}
				}
				fmt.Println(result)
			}
		}
	}

	return c.JSON(http.StatusOK, result)
}

func (h *Handler) RecivedInvoicesList(c echo.Context) error {
	req := new(model.ListReceivedReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	result, err := h.invoiceRepo.RecievedInvoiceList(req)
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func (h *Handler) RecievedInvoiceReject(c echo.Context) error {
	req := new(model.EtaInvoiceRejectBody)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	id := c.Param("id")
	d, err := utils.EtaRecentDocumentReject(&id, req)
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, d)
}
func (h *Handler) RecievedInvoiceListItems(c echo.Context) error {
	req := new(model.ListReceivedReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	result, err := h.invoiceRepo.RecievedInvoiceListItems(c.Param("id"))
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func (h *Handler) InvoicesRecentView(c echo.Context) error {
	d, err := utils.EtaRecentDocumentView(c.Param("id"))
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusOK, err.Error())
	}

	return c.JSON(http.StatusOK, d)
}

func (h *Handler) InvoicesList(c echo.Context) error {
	req := new(model.ListInvoicessRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	result, err := h.invoiceRepo.ListEInvoices(req)
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func (h *Handler) InvoicePost(c echo.Context) error {
	req := new(model.PostInvoicessRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	invoices, err := h.invoiceRepo.FindInvoiceData(req, h.companyInfo)
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// _, err = utils.SignInvoices(invoices)
	// if err := c.Bind(req); err != nil {
	// 	return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	// }
	return c.JSON(http.StatusOK, invoices)
}
