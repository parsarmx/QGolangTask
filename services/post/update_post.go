package post

import (
	"context"
	"encoding/json"
	"fmt"
	"task/models"
	"time"
)

func (postService *Service) UpdatePost(ctx context.Context, key string, data models.PostRequest) (bool, error) {

	exp := time.Duration(600 * time.Second) // 10 minutes

	value, err := json.Marshal(&data)
	if err != nil {
		fmt.Printf("(UpdatePost): %v", err)
		return false, err
	}

	err = postService.Client.Set(ctx, key, value, exp).Err()
	if err != nil {
		fmt.Printf("(UpdatePost): %v", err)
		return false, err
	}
	return true, nil

}
