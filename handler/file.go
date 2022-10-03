package handler

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Upload(c echo.Context) error {
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error getting file : "+err.Error())
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error opening file : "+err.Error())
	}
	defer src.Close()

	// Destination
	name := "assets/" + file.Filename
	dst, err := os.Create(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error creating file : "+err.Error())
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, "error copying file : "+err.Error())
	}

	return c.JSON(http.StatusOK, name)
}
