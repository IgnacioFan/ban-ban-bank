package wireset

import (
	"go-bank-express/internal/delivery/handler"

	"github.com/google/wire"
)

var HandlerSet = wire.NewSet(
	handler.NewPing,
	handler.NewUser,
)
