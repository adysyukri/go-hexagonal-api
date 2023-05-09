package handler

import (
	"go-api/internal/db/repo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Minimum of 100 deposit is required
type createAccountRequest struct {
	Deposit float64 `json:"deposit" binding:"required,min=100"`
}

func (h *Handler) createAccount(c *gin.Context) {
	userIDstring := c.Params.ByName("userID")

	userID, err := strconv.Atoi(userIDstring)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorMessageResponse(err))
		return
	}

	var req createAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorMessageResponse(err))
		return
	}

	p := &repo.AddAccountParams{
		UserID:  userID,
		Deposit: req.Deposit,
	}

	account, err := h.s.CreateAccount(c, p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorMessageResponse(err))
		return
	}

	c.JSON(http.StatusOK, account)
}

type createAccountNewUserRequest struct {
	Name    string  `json:"name" binding:"required"`
	Deposit float64 `json:"deposit" binding:"required,min=100"`
}

func (h *Handler) createAccountNewUser(c *gin.Context) {
	var req createAccountNewUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorMessageResponse(err))
		return
	}

	p := &repo.AddAccountNewUserParams{
		Name:    req.Name,
		Deposit: req.Deposit,
	}

	account, err := h.s.CreateAccountNewUser(c, p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorMessageResponse(err))
		return
	}

	c.JSON(http.StatusOK, account)
}

func (h *Handler) getBalance(c *gin.Context) {
	accountNumber := c.Params.ByName("accountNumber")

	p := &repo.GetAccountParams{
		AccountNumber: accountNumber,
	}

	balance, err := h.s.GetAccountBalance(c, p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorMessageResponse(err))
		return
	}

	c.JSON(http.StatusOK, balance)
}
