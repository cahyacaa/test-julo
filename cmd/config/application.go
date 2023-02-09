package config

type Application struct {
	Port string `env:"APP_PORT"`
	Host string `env:"APP_HOST"`
}
