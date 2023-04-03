// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package http

import (
	"go-bank-express/internal/conn"
	"go-bank-express/internal/delivery/handler"
	"go-bank-express/internal/delivery/http"
	"go-bank-express/internal/repository"
	"go-bank-express/internal/usecase"
)

// Injectors from wire.go:

func Initialize() (Application, error) {
	ping := handler.NewPing()
	db, err := conn.NewDatabase()
	if err != nil {
		return Application{}, err
	}
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUser(userUsecase)
	httpServer := http.NewHttpServer(ping, userHandler)
	application := newApplication(httpServer)
	return application, nil
}