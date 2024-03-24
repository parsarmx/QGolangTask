package selectors

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type PostSelector struct {
	Client *redis.Client
}

func NewPostSelector(client *redis.Client) *PostSelector {
	return &PostSelector{Client: client}
}

func (postSelector *PostSelector) IsSlugExists(ctx context.Context, key string) (int64, error) {

	val, err := postSelector.Client.Exists(ctx, key).Result()
	if err != nil {
		return val, err
	}
	return val, nil
}
