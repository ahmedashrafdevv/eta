package handler

import (
	"eta/utils"
	"net/http"
	"net/smtp"

	"github.com/labstack/echo/v4"
)

func (h *Handler) SendEmail(c echo.Context) error {

	// Sender data.
	from := "from@gmail.com"
	password := "<Email Password>"

	// Receiver email address.
	to := []string{
		"ahmed.ashraf.devv@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, "Email Sent Successfully!")

}
