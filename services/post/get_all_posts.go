package post

import (
	"context"
	"fmt"
	"task/models"
)

func (postService *Service) GetAllPosts(ctx context.Context) []models.AllPost {

	keys, err := postService.Client.Keys(ctx, "*").Result()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var posts []models.AllPost

	for _, key := range keys {
		value, err := postService.Client.Get(ctx, key).Result()
		if err != nil {
			fmt.Println(err)
			return nil
		}
		posts = append(posts, models.AllPost{Key: key, Data: value})
	}

	return posts
}
