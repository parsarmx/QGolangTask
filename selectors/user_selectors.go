package selectors

import (
	"task/models"

	"gorm.io/gorm"
)

type UserSelector struct {
	DB *gorm.DB
}

func NewUserSelector(db *gorm.DB) *UserSelector {
	return &UserSelector{DB: db}
}

func (UserSelector *UserSelector) GetUserByUsername(user *models.User, username string) {
	UserSelector.DB.Where("username = ?", username).Find(user)
}
