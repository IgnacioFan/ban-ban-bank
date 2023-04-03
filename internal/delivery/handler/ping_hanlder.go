package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingHandler struct{}

func NewPing() *PingHandler {
	return &PingHandler{}
}

func (p *PingHandler) Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Pong")
}
