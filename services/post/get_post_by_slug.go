package post

import (
	"context"
	"encoding/json"
	"fmt"
	"task/models"
)

func (postService *Service) GetPostBySlug(ctx context.Context, key string) models.Post {
	res, err := postService.Client.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return models.Post{}
	}

	var p models.Post
	err = json.Unmarshal([]byte(res), &p)

	return p

}
