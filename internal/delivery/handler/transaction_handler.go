package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct{}

func NewTransaction() *TransactionHandler {
	return &TransactionHandler{}
}

func (h *TransactionHandler) Transfer(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "start transfering")
}

func (h *TransactionHandler) ListTransactions(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "list transactions")
}
