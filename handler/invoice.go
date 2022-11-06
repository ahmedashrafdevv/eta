package handler

import (
	"eta/model"
	"eta/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func handleDate(s *string) {
	if idx := strings.Index(*s, "."); idx != -1 {
		*s = (*s)[:idx]
	}
	*s = strings.Replace(*s, "T", " ", 1)
	*s = strings.Replace(*s, "Z", "", 1)
}
func (h *Handler) InvoicesRecentEtl(c echo.Context) error {
	d, err := utils.EtaRecentDocuments()
	var result []model.EtaRecentDocumentsItem
	// var req []model.ReceivedInvoiceInsertReq
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusOK, err.Error())
	}
	for _, v := range d.Result {
		if v.IssuerId != "288271998" {
			existed := h.invoiceRepo.RecievedInvoiceCheckExists(&v.Uuid)
			if !existed {
				v.TotalTax = v.Total - v.TotalSales
				handleDate(&v.DateTimeIssued)
				handleDate(&v.DateTimeReceived)

				v.TotalTax = utils.RoundFloat(&v.TotalTax, 5)

				invoiceDetails, err := utils.EtaRecentDocumentView(v.Uuid)
				if utils.CheckErr(&err) {
					return c.JSON(http.StatusOK, err.Error())
				}

				_, err = h.invoiceRepo.RecievedInvoiceInsert(v, invoiceDetails.InvoiceLines)
				if utils.CheckErr(&err) {
					return c.JSON(http.StatusOK, err.Error())
				}
				result = append(result, v)
			}
		}
	}

	return c.JSON(http.StatusOK, d)
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
