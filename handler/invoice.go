package handler

import (
	"eta/model"
	"eta/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) InvoicesRecentList(c echo.Context) error {
	d, err := utils.EtaRecentDocuments()
	var result []model.EtaRecentDocumentsItem
	// var req []model.ReceivedInvoiceInsertReq
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusOK, err.Error())
	}
	for _, v := range d.Result {
		if v.IssuerId != "288271998" {
			v.TotalTax = v.Total - v.TotalSales
			v.TotalTax = utils.RoundFloat(&v.TotalTax, 5)
			result = append(result, v)
			invoiceDetails, err := utils.EtaRecentDocumentView(v.Uuid)
			if utils.CheckErr(&err) {
				return c.JSON(http.StatusOK, err.Error())
			}
			return c.JSON(http.StatusOK, invoiceDetails)

		}

	}

	// var req model.ReceivedInvoiceInsertReq

	// req.Invoice.InternalId = "Y-2113-202200017451-2022102578000061264"
	// req.Invoice.TotalAmount = 1546.18
	// req.Invoice.TotalTax = 1546.18 - 1356.3
	// req.Invoice.IssuerName = "Al-Futtaim for Commercial and Administrative Centres S.A.E."
	// req.Invoice.IssuerRin = "339745703"
	// req.Invoice.DateTimeIssued = "2022-10-25 00:00:00"
	// req.Invoice.DateTimeRecieved = "2022-10-25 18:37:00"

	// var item1 model.ReceivedInvoiceItem
	// item1.ItemName = "Temporary Classification \n التصنيف المؤقت \n Chilled water-recoveries -Cairo-From 23/09/2022 To 22/10/2022"
	// item1.ItemType = "EGS"
	// item1.ItemCode = "99999999"
	// item1.Price = 1356.3
	// item1.Quantity = 1.0
	// item1.TotalAmount = 1546.18
	// item1.TotalTax = 1546.18 - 1356.3
	// item1.SubTotal = 1356.3

	// req.Items = append(req.Items, item1)

	// _, err = h.invoiceRepo.RecievedInvoiceInsert(&req)
	// if utils.CheckErr(&err) {
	// 	return c.JSON(http.StatusOK, err.Error())
	// }
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
