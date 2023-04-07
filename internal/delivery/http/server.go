package http

import (
	"go-bank-express/internal/delivery/handler"
	"go-bank-express/internal/delivery/http/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	*gin.Engine
	PingHandler   *handler.PingHandler
	UserHandler   *handler.UserHandler
	WalletHandler *handler.WalletHandler
}

func NewHttpServer(ping *handler.PingHandler, user *handler.UserHandler, wallet *handler.WalletHandler) *HttpServer {
	server := &HttpServer{
		Engine:        gin.Default(),
		PingHandler:   ping,
		UserHandler:   user,
		WalletHandler: wallet,
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
		v1.POST("/register", s.UserHandler.RegisterUser)
		v1.POST("/login", s.UserHandler.LoginUser)
		v1.Use(middlewares.Auth)
		v1.POST("/wallet", s.WalletHandler.CreateWallet)
		v1.GET("/wallets", s.WalletHandler.ListWallets)
		v1.POST("/wallet/balance", s.WalletHandler.UpdateWalletBalance)
	}
}
