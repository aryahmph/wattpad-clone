package env

import (
	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	Port          int    `env:"PORT,unset"`
	RedisAddress  string `env:"REDIS_ADDRESS,unset"`
	RedisUsername string `env:"REDIS_USERNAME,unset"`
	RedisPassword string `env:"REDIS_PASSWORD,unset"`
	KaguyaAddress string `env:"KAGUYA_ADDRESS,unset"`
}

func LoadConfig() *config {
	cfg := new(config)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
