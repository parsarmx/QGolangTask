package configs

import (
	"os"
	"strconv"
)

type AuthConfig struct {
	AccessSecret  string
	RefreshSecret string
	AcessTTL      int
	RefreshTTL    int
}

func LoadAuthConfig() AuthConfig {

	a_ttl, _ := strconv.Atoi(os.Getenv("ACCESS_TTL"))
	r_ttl, _ := strconv.Atoi(os.Getenv("REFRESH_TTL"))

	return AuthConfig{
		AccessSecret:  os.Getenv("ACCESS_SECRET"),
		RefreshSecret: os.Getenv("REFRESH_SECRET"),
		AcessTTL:      a_ttl,
		RefreshTTL:    r_ttl,
	}
}
