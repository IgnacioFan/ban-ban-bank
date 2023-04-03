package wireset

import (
	"go-bank-express/internal/usecase"

	"github.com/google/wire"
)

var UsecaseSet = wire.NewSet(
	usecase.NewUserUsecase,
)
