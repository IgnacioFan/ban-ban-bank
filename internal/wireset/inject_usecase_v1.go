package wireset

import (
	v1 "go-bank-express/internal/usecase/v1"

	"github.com/google/wire"
)

var UsecaseV1Set = wire.NewSet(
	v1.NewUserUsecase,
)
