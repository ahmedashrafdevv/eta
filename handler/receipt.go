package handler

import (
	"crypto/sha256"
	"encoding/hex"
	"eta/model"
	"eta/utils"
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
	var recieptsReq model.ReceiptSubmitRequest
	receipts, err := h.receiptRepo.FindUnPostedReciepts(h.companyInfo)
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	for i := 0; i < len(receipts); i++ {
		serialized := utils.SerializeInvoice(receipts[i])
		hash := sha256.New()
		hash.Write([]byte(serialized))
		sha1_hash := hex.EncodeToString(hash.Sum(nil))
		receipts[i].Header.Uuid = sha1_hash
		recieptsReq.Receipts = append(recieptsReq.Receipts, receipts[i])
	}

	resp, err := utils.SubmitReceipt(&recieptsReq)
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
