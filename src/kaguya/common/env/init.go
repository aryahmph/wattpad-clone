package env

import (
	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	Port                  int    `env:"PORT,unset"`
	Host                  string `env:"HOST,unset"`
	PostgresMigrationPath string `env:"POSTGRES_MIGRATION_PATH,unset"`
	PostgresURL           string `env:"POSTGRES_CONNECTION_URL,unset"`
	AccessTokenKey        string `env:"ACCESS_TOKEN_KEY,unset"`
	NSQAddr               string `env:"NSQ_ADDR,unset"`
	NSQMailTopic          string `env:"NSQ_MAIL_TOPIC,unset"`
}

func LoadConfig() *config {
	cfg := new(config)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
