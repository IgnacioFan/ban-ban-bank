package v1

import (
	"go-bank-express/internal/entity"
	"go-bank-express/internal/repository"
	"go-bank-express/internal/usecase"
)

type UserUsecase struct {
	UserRepo *repository.UserRepository
}

func NewUserUsecase(repo *repository.UserRepository) usecase.UserUsecase {
	return &UserUsecase{
		UserRepo: repo,
	}
}

func (u *UserUsecase) Create() error {
	user := &entity.User{
		Name: "Foo Bar",
	}
	if err := u.UserRepo.CreateUser(user); err != nil {
		return err
	}
	return nil
}
