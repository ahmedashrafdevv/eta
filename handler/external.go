package handler

import (
	"eta/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) EtaLogin(c echo.Context) error {
	var req model.EtaLoginRequest
	fmt.Println(req.ClientId)
	return c.JSON(http.StatusOK, "Email Sent Successfully!")

}
