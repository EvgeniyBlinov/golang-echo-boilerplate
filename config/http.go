package config

import "os"

type HTTPConfig struct {
	Host         string
	Port         string
	ExposePort   string
	AllowOrigins string
}

func LoadHTTPConfig() HTTPConfig {
	return HTTPConfig{
		Host:         os.Getenv("HOST"),
		Port:         os.Getenv("PORT"),
		ExposePort:   os.Getenv("EXPOSE_PORT"),
		AllowOrigins: os.Getenv("ALLOW_ORIGINS"),
	}
}
