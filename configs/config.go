package configs

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	AUTH  AuthConfig
	DB    DBConfig
	HTTP  HttpConfig
	Redis RedisConfig
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		AUTH:  LoadAuthConfig(),
		DB:    LoadDBConfig(),
		HTTP:  LoadHTTPConfig(),
		Redis: LoadRedisConfig(),
	}
}
