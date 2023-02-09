package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

var GlobalConfig Config

const path string = "cmd/config/.env"

type Config struct {
	App   Application
	Cache Cache
}

func InitializeConfig() {
	var cfg Config

	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	GlobalConfig = cfg
}
