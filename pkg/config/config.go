package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP HTTP
		DB   DB
		// Services
		Services struct {
			User User
		}
	}
	DB struct {
		Host     string `env:"DB_HOST"`
		Port     int    `env:"DB_PORT"`
		User     string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
		Name     string `env:"DB_NAME"`
		SSLMode  string `env:"DB_SSLMODE"`
	}

	User struct {
		Port string `env:"USER_PORT"`
	}

	HTTP struct {
		Port string `env:"HTTP_PORT"`
	}
)

func New() (*Config, error) {
	cfg := Config{}
	err := cleanenv.ReadEnv(&cfg)
	return &cfg, err
}
