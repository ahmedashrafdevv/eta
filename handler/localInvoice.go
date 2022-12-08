package handler

import (
	"eta/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) AccountsList(c echo.Context) error {
	req := new(model.GetAccountRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	accounts, err := h.localInvoiceRepo.AccountsList(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, accounts)
}

func (h *Handler) InvoicesLocalList(c echo.Context) error {
	// code, _ := strconv.Atoi(string(c.Get("id")))

	req := new(model.ListInvoiceReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	orders, err := h.localInvoiceRepo.ELocalInvoicesList(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, orders)
}

func (h *Handler) InvoicesLocalClose(c echo.Context) error {
	code, _ := strconv.Atoi(c.Param("id"))
	items, err := h.localInvoiceRepo.ELocalInvoicesOrderClose(code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, items)
}
func (h *Handler) InvoicesLocalUpdateItem(c echo.Context) error {
	req := new(model.UpdateOrderItemReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	items, err := h.localInvoiceRepo.ELocalInvoicesOrderItemUpdate(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, items)
}
func (h *Handler) InvoicesLocalListItems(c echo.Context) error {
	code, _ := strconv.Atoi(c.Param("id"))
	items, err := h.localInvoiceRepo.ELocalInvoicesListItems(code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, items)
}

func (h *Handler) InvoicesLocalDeleteItem(c echo.Context) error {
	code, _ := strconv.Atoi(c.Param("id"))
	items, err := h.localInvoiceRepo.ELocalInvoicesOrderItemDelete(code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, items)
}
func (h *Handler) InvoicesLocalDocNo(c echo.Context) error {
	// code, _ := strconv.Atoi(string(c.Get("id")))
	transSerial, _ := strconv.Atoi(c.QueryParam("transSerial"))
	resp, err := h.localInvoiceRepo.ELocalInvoicesDocNo(transSerial)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	currentDocNo := *resp + 1
	return c.JSON(http.StatusOK, currentDocNo)
}

func (h *Handler) InvoicesLocalOrderInsert(c echo.Context) error {
	// code := c.Get("id").(int)
	fmt.Println("c.Get()")
	fmt.Println(c.Get("exp"))
	// fmt.Println(c.Get())
	req := new(model.InsertOrderReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	// req.EmpCode = code
	id, err := h.localInvoiceRepo.ELocalInvoicesOrderInsert(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, id)
}

func (h *Handler) InvoicesLocalOrderItemInsert(c echo.Context) error {
	req := new(model.InsertOrderItemReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	id, err := h.localInvoiceRepo.ELocalInvoicesOrderItemInsert(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, id)
}

func (h *Handler) ItemsList(c echo.Context) error {
	store, _ := strconv.Atoi(c.QueryParam("store"))
	resp, err := h.localInvoiceRepo.ListItems(store)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, resp)
}
