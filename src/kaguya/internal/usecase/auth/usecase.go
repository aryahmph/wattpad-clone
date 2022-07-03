package auth

import (
	"context"
	userDomain "github.com/aryahmph/wattpad-clone/src/kaguya/internal/domain/user"
)

type Usecase interface {
	Login(ctx context.Context, login userDomain.Login) (token string, err error)
	Session(ctx context.Context, token string) (user userDomain.User, err error)
}
