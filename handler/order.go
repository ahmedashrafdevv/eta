package handler

import (
	"eta/model"
	"eta/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) OrdersListByTransSerialStoreConvertedDate(c echo.Context) error {
	req := new(model.ListOrdersRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	result, err := h.orderRepo.ListByTransSerialStoreConvertedDate(req)
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func (h *Handler) OrdersConvertToEta(c echo.Context) error {
	serial, _ := strconv.ParseInt(c.Param("serial"), 0, 64)
	result, err := h.orderRepo.ConvertToEtaInvoice(&serial)
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
