package auth

import "context"

type Usecase interface {
	SetSession(ctx context.Context, token, data string) (err error)
	GetSession(ctx context.Context, token string) (data string, err error)
}
