package db

import (
	"context"
	"fmt"
	"task/configs"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *configs.Config) *gorm.DB {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Name,
		cfg.DB.Password)

	fmt.Println(dataSourceName)

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}

func InitRedis(cfg *configs.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       0,
	})

	// check if redis connected or not, it should print 'redis: PONG'
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("redis: %v", pong)

	return client

}
