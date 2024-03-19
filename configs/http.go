package configs

import "os"

type HttpConfig struct {
	Host string
	Port string
}

func LoadHTTPConfig() HttpConfig {
	return HttpConfig{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	}
}
