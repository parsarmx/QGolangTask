package post

import (
	"context"
	"encoding/json"
	"fmt"
	"task/models"
	"time"
)

func (postService *Service) CreatePost(key string, data models.PostRequest) (bool, error) {
	var ctx = context.Background()

	exp := time.Duration(600 * time.Second) // 10 minutes

	value, err := json.Marshal(&data)
	if err != nil {
		fmt.Printf("(CreatePost): %v", err)
		return false, err
	}

	err = postService.Client.Set(ctx, key, value, exp).Err()
	if err != nil {
		fmt.Printf("(CreatePost): %v", err)
		return false, err
	}

	return true, nil
}
