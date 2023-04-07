package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	_headerRequestID = "X-Request-Id"
)

func RequestID(ctx *gin.Context) {
	requestId := ctx.GetHeader(_headerRequestID)
	if len(requestId) == 0 {
		requestId = uuid.New().String()
		ctx.Request.Header.Set(_headerRequestID, requestId) // use ctx.Request.Header to get headers
	}
	ctx.Set(_headerRequestID, requestId) // use ctx.Keys to get key/value pairs

	ctx.Next()
}
