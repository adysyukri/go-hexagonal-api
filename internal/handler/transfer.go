package handler

import (
	"go-api/internal/db/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Minimum of 10 for the amount to transfer
type makeTransferRequest struct {
	FromAccountNumber string  `json:"from_account_number" binding:"required"`
	ToAccountNumber   string  `json:"to_account_number" binding:"required"`
	Amount            float64 `json:"amount" binding:"required,min=10"`
}

func (h *Handler) makeTransfer(c *gin.Context) {

	var req makeTransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorMessageResponse(err))
		return
	}

	p := &repo.ExecTransferParams{
		FromAccountNumber: req.FromAccountNumber,
		ToAccountNumber:   req.ToAccountNumber,
		Amount:            req.Amount,
	}

	transfer, err := h.s.MakeTransfer(c, p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorMessageResponse(err))
		return
	}

	c.JSON(http.StatusOK, transfer)
}

func (h *Handler) getTransfers(c *gin.Context) {
	accountNumber := c.Params.ByName("accountNumber")

	p := &repo.GetTransfersParams{
		AccountNumber: accountNumber,
	}

	list := h.s.ListTransfers(c, p)

	c.JSON(http.StatusOK, list)
}
