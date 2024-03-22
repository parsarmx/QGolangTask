package token

import "task/configs"

type Service struct {
	config *configs.Config
}

func NewTokenService(cfg *configs.Config) *Service {
	return &Service{
		config: cfg,
	}
}
