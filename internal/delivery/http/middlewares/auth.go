package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context) {
	fmt.Println("The endpoint needs to be passed from http authorization")
	ctx.Next()
}
