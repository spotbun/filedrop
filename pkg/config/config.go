package config

import "os"

type Config struct {
	Title string

	initiated bool
}

var config Config

func GetConfig() Config {
	if config.initiated {
		return config
	}

	var cfg Config
	cfg.Title = getFromEnv("TITLE", "Filedrop")

	cfg.initiated = true

	config = cfg
	return config
}

func getFromEnv(key, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return def
}
