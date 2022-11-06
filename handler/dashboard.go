package handler

import (
	"eta/model"
	"eta/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) DashboardStats(c echo.Context) error {

	result, err := h.orderRepo.DashboardStats()
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func (h *Handler) DashboardStoreStats(c echo.Context) error {
	req := new(model.DashboardStatsRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	result, err := h.orderRepo.DashboardStoreStats(req)
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
