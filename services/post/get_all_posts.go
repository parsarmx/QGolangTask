package post

import (
	"context"
	"fmt"
	"task/models"
)

func (postService *Service) GetAllPosts(ctx context.Context) []models.AllPost {

	slugs, err := postService.Client.Keys(ctx, "*").Result()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(slugs)
	var posts []models.AllPost

	for _, slug := range slugs {
		value, err := postService.Client.Get(ctx, slug).Result()
		if err != nil {
			fmt.Println(err)
			return nil
		}
		posts = append(posts, models.AllPost{Slug: slug, Data: value})
	}

	return posts
}
