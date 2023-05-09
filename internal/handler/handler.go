package handler

import (
	"go-api/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	s      services.Servicer
	router *gin.Engine
}

func NewHandler(s services.Servicer) *Handler {
	h := &Handler{s: s}
	router := gin.Default()

	router.POST("/account", h.createAccountNewUser)
	router.POST("/account/:userID", h.createAccount)
	router.GET("/account/:accountNumber", h.getBalance)
	router.POST("/transfer", h.makeTransfer)
	router.GET("/transfer/:accountNumber", h.getTransfers)

	h.router = router

	return h
}

func errorMessageResponse(err error) gin.H {
	return gin.H{"message": err.Error()}
}

func (h *Handler) HandlerRouter() *gin.Engine {
	return h.router
}
