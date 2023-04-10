package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	BaseHandler
}

func NewUser() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) RegisterUser(ctx *gin.Context) {
	// requestId := ctx.Keys["X-Request-Id"].(string)
	req := RegisterUserReq{}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// pass req and req_id to usecase
	err := h.PubMessage(ctx, &req)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, "error return")
		return
	}
	// data := h.Producer.Listen(requestId)
	// fmt.Println(data)
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
