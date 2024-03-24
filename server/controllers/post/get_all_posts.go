package post

import (
	"net/http"
	"task/responses"
	"task/services/post"

	"github.com/labstack/echo/v4"
)

func (postController *PostController) GetAllPosts(ctx echo.Context) error {
	postService := post.NewPostService(postController.server.Redis)
	data := postService.GetAllPosts(ctx.Request().Context())

	return responses.Response(ctx, http.StatusOK, data)
}
