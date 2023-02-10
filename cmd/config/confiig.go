package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

var GlobalConfig Config

const path string = "cmd/config/.env"

type Config struct {
	App   Application
	Cache Cache
}

func InitializeConfig(isProdEnv bool) {
	var cfg Config
	var err error
	if isProdEnv {
		err = cleanenv.ReadEnv(&cfg)
	} else {
		err = cleanenv.ReadConfig(path, &cfg)
	}
	fmt.Println(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	GlobalConfig = cfg
}
