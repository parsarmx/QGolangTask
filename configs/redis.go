package configs

import (
	"os"
)

type RedisConfig struct {
	Addr     string
	Password string
}

func LoadRedisConfig() RedisConfig {
	return RedisConfig{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}
}
