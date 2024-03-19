package db

import (
	"fmt"
	"task/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(cfg *configs.Config) *gorm.DB {
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
