package wireset

import (
	"go-bank-express/internal/repository"

	"github.com/google/wire"
)

var RepositorySet = wire.NewSet(
	repository.NewUserRepository,
)
