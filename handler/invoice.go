package handler

import (
	"eta/model"
	"eta/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

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
