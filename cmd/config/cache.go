package config

type Cache struct {
	DB                int    `env:"REDIS_DB"`
	Host              string `env:"REDIS_HOST"`
	Port              int    `env:"REDIS_PORT"`
	Username          string `env:"REDIS_USERNAME"`
	Password          string `env:"REDIS_PASSWORD"`
	IdempotentExpired int    `env:"REDIS_IDEMPOTENT_EXPIRED"`
}
