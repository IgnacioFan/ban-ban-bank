package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
}

func NewWallet() *WalletHandler {
	return &WalletHandler{}
}

func (h *WalletHandler) CreateWallet(ctx *gin.Context) {
	req := CreateWalletReq{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// pass req and req_id to usecase
	ctx.JSON(http.StatusOK, fmt.Sprintf("created wallet: %v", req))
}

func (h *WalletHandler) ListWallets(ctx *gin.Context) {
	req := ListWalletReq{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// pass req and req_id to usecase
	ctx.JSON(http.StatusOK, fmt.Sprintf("list wallets: %v", req))
}

func (h *WalletHandler) UpdateWalletBalance(ctx *gin.Context) {
	req := UpdateWalletBalanceReq{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// pass req and req_id to usecase
	ctx.JSON(http.StatusOK, fmt.Sprintf("update wallet balance: %v", req))
}
