package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
}

func NewWallet() *WalletHandler {
	return &WalletHandler{}
}

func (h *WalletHandler) CreateWallet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "created a new wallet")
}
