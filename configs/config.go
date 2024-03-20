package configs

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
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
		DB:    LoadDBConfig(),
		HTTP:  LoadHTTPConfig(),
		Redis: LoadRedisConfig(),
	}
}
