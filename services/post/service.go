package post

import "github.com/redis/go-redis/v9"

type Service struct {
	Client *redis.Client
}

func NewPostService(client *redis.Client) *Service {
	return &Service{Client: client}
}
