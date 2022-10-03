package handler

import (
	"eta/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) StoresListAll(c echo.Context) error {
	result, err := h.storeRepo.ListAll()
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
