package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct{}

func NewTransaction() *TransactionHandler {
	return &TransactionHandler{}
}

func (h *TransactionHandler) Transfer(ctx *gin.Context) {
	req := TransferReq{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// sent req and req_id to usecase
	ctx.JSON(http.StatusOK, fmt.Sprintf("transfer: %v", req))
}

func (h *TransactionHandler) ListTransactions(ctx *gin.Context) {
	req := ListTransactionsReq{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// sent req and req_id to usecase
	ctx.JSON(http.StatusOK, fmt.Sprintf("list transactions: %v", req))
}
