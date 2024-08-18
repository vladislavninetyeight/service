package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Server Server `yaml:"server"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func MustLoad() *Config {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("CONFIG_PATH does not exist: %s", path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		log.Fatalf("CONFIG could not be loaded: %s", err)
	}

	return &cfg
}
