//go:build wireinject
// +build wireinject

package http

import (
	"go-bank-express/internal/conn"
	"go-bank-express/internal/delivery/http"
	"go-bank-express/internal/wireset"

	"github.com/google/wire"
)

func Initialize() (Application, error) {
	wire.Build(
		newApplication,
		http.NewHttpServer,
		conn.NewDatabase,
		wireset.HandlerSet,
		wireset.UsecaseSet,
		wireset.RepositorySet,
	)
	return Application{}, nil
}
