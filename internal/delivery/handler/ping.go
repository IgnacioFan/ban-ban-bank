package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Ping struct{}

func NewPing() *Ping {
	return &Ping{}
}

func (p *Ping) Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Pong")
}
