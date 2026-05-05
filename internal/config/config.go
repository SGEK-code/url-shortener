package config

import (
	"flag"
	"os"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	ListenAddr string `env:"SERVER_ADDRESS"`
	BaseURL    string `env:"BASE_URL"`
	LogLevel   string `env:"LOG_LEVEL"`
}

func ParseConfig() *Config {
	defaults := Config{
		ListenAddr: ":8080",
		BaseURL:    "http://localhost:8080",
		LogLevel:   "info",
	}
	cfg := defaults

	envSet := map[string]bool{
		"SERVER_ADDRESS": envIsSet("SERVER_ADDRESS"),
		"BASE_URL":       envIsSet("BASE_URL"),
		"LOG_LEVEL":      envIsSet("LOG_LEVEL"),
	}

	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	var (
		listenAddr = flag.String("a", defaults.ListenAddr, "Address to listen on (e.g. localhost:8888)")
		baseURL    = flag.String("b", defaults.BaseURL, "Base URL for shortened links (e.g. http://localhost:8000)")
		logLevel   = flag.String("l", defaults.LogLevel, "Logging level (e.g. debug, info...)")
	)

	flag.Parse()

	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "a":
			if !envSet["SERVER_ADDRESS"] {
				cfg.ListenAddr = *listenAddr
			}
		case "b":
			if !envSet["BASE_URL"] {
				cfg.BaseURL = *baseURL
			}
		case "l":
			if !envSet["LOG_LEVEL"] {
				cfg.LogLevel = *logLevel
			}
		}
	})

	return &cfg
}

// envIsSet проверяет, существует ли переменная окружения
func envIsSet(key string) bool {
	_, exists := os.LookupEnv(key)
	return exists
}
