package auth

import (
	"context"
	"time"
)

type Repository interface {
	Set(ctx context.Context, key string, data any, duration time.Duration) (err error)
	Get(ctx context.Context, key string) (data string, err error)
}
