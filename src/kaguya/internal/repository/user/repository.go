package user

import (
	"context"
	userDomain "github.com/aryahmph/wattpad-clone/src/kaguya/internal/domain/user"
)

type Repository interface {
	Insert(ctx context.Context, user userDomain.User) (rid string, err error)
	FindByID(ctx context.Context, id string) (user userDomain.User, err error)
	FindByUsername(ctx context.Context, username string) (user userDomain.User, err error)
	FindByEmail(ctx context.Context, email string) (user userDomain.User, err error)
	UpdatePassword(ctx context.Context, id string, password string) (rid string, err error)
}
