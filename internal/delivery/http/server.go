package http

import (
	"go-bank-express/internal/delivery/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	*gin.Engine
	PingHandler *handler.PingHandler
	UserHandler *handler.UserHandler
}

func NewHttpServer(ping *handler.PingHandler, user *handler.UserHandler) *HttpServer {
	server := &HttpServer{
		Engine:      gin.Default(),
		PingHandler: ping,
		UserHandler: user,
	}
	server.SetRouter()
	return server
}

func (s *HttpServer) SetRouter() {
	s.Engine.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "Not Found")
	})

	v1 := s.Engine.Group("/api/v1")

	{
		v1.GET("/ping", s.PingHandler.Pong)
		v1.GET("/user", s.UserHandler.CreateUser)
	}
}
