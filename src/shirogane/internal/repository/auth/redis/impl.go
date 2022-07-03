package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type redisAuthRepositoryImpl struct {
	client *redis.Client
}

func NewRedisAuthRepositoryImpl(client *redis.Client) redisAuthRepositoryImpl {
	return redisAuthRepositoryImpl{client: client}
}

func (repository redisAuthRepositoryImpl) Set(ctx context.Context, key string, data any, duration time.Duration) (err error) {
	err = repository.client.Set(ctx, key, data, duration).Err()
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (repository redisAuthRepositoryImpl) Get(ctx context.Context, key string) (data string, err error) {
	err = repository.client.Get(ctx, key).Scan(&data)
	switch {
	case err == redis.Nil:
		return "", status.Error(codes.NotFound, err.Error())
	case err != nil:
		return "", status.Error(codes.Internal, err.Error())
	}
	return data, nil
}
