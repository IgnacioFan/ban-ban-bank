package handler

import (
	"fmt"
	"go-bank-express/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

func NewUser(user usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: user,
	}
}

func (h *UserHandler) RegisterUser(ctx *gin.Context) {
	req := RegisterUserReq{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// pass req and req_id to usecase
	// if err := h.UserUsecase.Create(); err != nil {
	// 	ctx.JSON(http.StatusNotFound, err.Error())
	// }
	ctx.JSON(http.StatusOK, fmt.Sprintf("register user %s", req))
}

func (h *UserHandler) LoginUser(ctx *gin.Context) {
	req := LoginUserReq{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	req.IP = ctx.ClientIP()
	// pass req and req_id to usecase
	ctx.JSON(http.StatusOK, fmt.Sprintf("login user %s", req))
}
