package handler

import (
	"go-bank-express/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase *usecase.UserUsecase
}

func NewUser(user *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: user,
	}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	if err := h.UserUsecase.Create(); err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
	}
	ctx.JSON(http.StatusOK, "created user")
}
