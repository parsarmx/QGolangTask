package user

import (
	"task/models"

	"golang.org/x/crypto/bcrypt"
)

func (userService *Service) RegisterUser(req models.RegisterRequest) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Username: req.Username,
		Password: string(hashedPassword),
	}

	err = userService.DB.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
