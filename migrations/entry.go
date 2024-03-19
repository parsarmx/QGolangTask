package migrations

import (
	"fmt"
	"log"
	"sync"
	"task/configs"
	"task/models/account"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var onceDB sync.Once
var DB *gorm.DB

func Migrate(db *configs.DBConfig) {
	onceDB.Do(func() {
		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN: fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				db.Host,
				db.Port,
				db.User,
				db.Password,
				db.Name,
			),
		}), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
		db.AutoMigrate(&account.User{})

		DB = db
	})
}
