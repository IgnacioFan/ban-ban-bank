package app

import (
	"fmt"
	"go-bank-express/internal/delivery/http"
)

type Application struct {
	httpServer *http.HttpServer
}

func (app Application) Start(port int) error {
	return app.httpServer.Run(fmt.Sprintf(":%d", port))
}

func newApplication(
	httpServer *http.HttpServer,
) Application {
	return Application{
		httpServer: httpServer,
	}
}
