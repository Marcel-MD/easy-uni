package config

import (
	"sync"
	"time"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Port        string `env:"PORT" envDefault:":8080"`
	AllowOrigin string `env:"ALLOW_ORIGIN" envDefault:"*"`

	ApiSecret     string        `env:"API_SECRET" envDefault:"SecretSecretSecret"`
	TokenLifespan time.Duration `env:"TOKEN_LIFESPAN" envDefault:"24h"`

	DatabaseUrl string `env:"DATABASE_URL" envDefault:"postgres://postgres:password@postgres:5432/users"`

	SmtpHost     string `env:"SMTP_HOST" envDefault:"smtp.gmail.com"`
	SmtpPort     string `env:"SMTP_PORT" envDefault:"587"`
	SmtpEmail    string `env:"SMTP_EMAIL" envDefault:""`
	SmtpPassword string `env:"SMTP_PASSWORD" envDefault:""`
	SenderName   string `env:"SENDER_NAME" envDefault:""`
}

var (
	once sync.Once
	cfg  Config
)

func GetConfig() Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Warn().Err(err).Msg("Failed to load .env file.")
		}

		if err := env.Parse(&cfg); err != nil {
			log.Fatal().Err(err).Msg("Failed to parse environment variables.")
		}
	})

	return cfg
}
