package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

func NewRedisClient(addr, username, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: username,
		Password: password,
		DB:       0,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatalln(err)
	}
	return client
}
