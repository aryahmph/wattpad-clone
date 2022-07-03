package env

import (
	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	MailEmail     string `env:"MAIL_EMAIL,unset"`
	MailPassword  string `env:"MAIL_PASSWORD"`
	MailSMTPHost  string `env:"MAIL_SMTP_HOST,unset"`
	MailSMTPPort  int    `env:"MAIL_SMTP_PORT,unset"`
	NSQTopic      string `env:"NSQ_TOPIC,unset"`
	NSQChannel    string `env:"NSQ_CHANNEL,unset"`
	NSQLookupAddr string `env:"NSQ_LOOKUP_ADDR,unset"`
}

func LoadConfig() *config {
	cfg := new(config)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
