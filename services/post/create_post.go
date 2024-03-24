package post

import (
	"context"
	"encoding/json"
	"task/models"
	"time"
)

func (postService *Service) CreatePost(key string, data models.PostRequest) (bool, error) {
	var ctx = context.Background()

	exp := time.Duration(600 * time.Second) // 10 minutes
	value, err := json.Marshal(data)
	if err != nil {
		return false, err
	}
	postService.Client.Set(ctx, key, string(value), exp)

	return true, nil
}
