package token

import (
	"log"
	"task/models"
	"time"

	"github.com/golang-jwt/jwt"
)

func (tokenService *Service) GenerateToken(username string, userID uint) (string, string, error) {
	//ACCESS
	accessClaims := models.CustomClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			// this should be store in .env
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(tokenService.config.AUTH.AcessTTL)).Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessSignedToken, err := accessToken.SignedString(models.JWTSecret)
	if err != nil {
		log.Println("(GenerateAccessToken) Error :", err)
		return "", "", err
	}

	// REFRESH
	refreshClaims := models.CustomClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(tokenService.config.AUTH.RefreshTTL)).Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshSignedToken, err := refreshToken.SignedString(models.JWTSecret)
	if err != nil {
		log.Println("(GenerateRefreshToken) Error :", err)
		return "", "", err
	}

	return accessSignedToken, refreshSignedToken, nil
}
