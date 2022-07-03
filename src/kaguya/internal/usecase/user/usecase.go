package user

import (
	"context"
	userDomain "github.com/aryahmph/wattpad-clone/src/kaguya/internal/domain/user"
)

type Usecase interface {
	Register(ctx context.Context, registrar userDomain.Register) (userDomain.User, error)
}
