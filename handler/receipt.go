package handler

import (
	"crypto/sha256"
	"encoding/hex"
	"eta/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) ReceiptsListByPosted(c echo.Context) error {
	// resp, err := utils.SubmitInvoice()
	// if utils.CheckErr(&err) {
	// 	return c.JSON(http.StatusOK, err.Error())
	// }
	// // postedFilter, _ := strconv.ParseBool(c.QueryParam("posted"))
	// // result, err := h.receiptRepo.ListReceiptsByPosted(&postedFilter)
	// // utils.CheckErr(&err)
	return c.JSON(http.StatusOK, "resp")
}

func (h *Handler) ReceiptPost(c echo.Context) error {

	receipts, err := h.receiptRepo.FindUnPostedReciepts(h.companyInfo)
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusOK, err.Error())
	}

	serialized := utils.SerializeInvoice(receipts[0])
	hash := sha256.New()
	hash.Write([]byte(serialized))
	sha1_hash := hex.EncodeToString(hash.Sum(nil))
	receipts[0].Header.Uuid = sha1_hash
	fmt.Println("serialized")
	fmt.Println(serialized)
	return c.JSON(http.StatusOK, receipts[0])
}
