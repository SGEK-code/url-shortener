package config

import (
	"flag"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	ListenAddr string `env:"SERVER_ADDRESS"` // флаг -a
	BaseURL    string `env:"BASE_URL"`       // флаг -b
	LogLevel   string `env:"LOG_LEVEL"`      // флаг -l
}

func ParseConfig() *Config {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}

	if cfg.ListenAddr == "" {
		flag.StringVar(&cfg.ListenAddr, "a", ":8080", "Address to listen on (e.g. localhost:8888)")
	}
	if cfg.BaseURL == "" {
		flag.StringVar(&cfg.BaseURL, "b", "http://localhost:8080",
			"Base URL for shortened links (e.g. http://localhost:8000)")
	}
	if cfg.LogLevel == "" {
		flag.StringVar(&cfg.LogLevel, "l", "info",
			"Logging level (e.g. debug, info...)")
	}

	flag.Parse()

	return &cfg
}
