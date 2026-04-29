package config

import (
	"flag"
)

type Config struct {
	ListenAddr string // флаг -a
	BaseURL    string // флаг -b
}

func ParseConfig() *Config {
	var cfg Config
	flag.StringVar(&cfg.ListenAddr, "a", ":8080", "Address to listen on (e.g. localhost:8888)")
	flag.StringVar(&cfg.BaseURL, "b", "http://localhost:8080", "Base URL for shortened links (e.g. http://localhost:8000)")

	flag.Parse()

	return &cfg
}
