package post

import (
	"fmt"
	"net/http"
	"task/responses"
	"task/selectors"
	"task/services/post"

	"github.com/labstack/echo/v4"
)

func (postController *PostController) GetPostBySlug(ctx echo.Context) error {
	slug := ctx.Param("slug")

	key := fmt.Sprintf("post:%v", slug)

	postSelector := selectors.NewPostSelector(postController.server.Redis)
	val, _ := postSelector.IsSlugExists(ctx.Request().Context(), key)
	if val != 1 {
		return responses.ErrorResponse(ctx, http.StatusBadRequest, responses.SLUG_DOES_NOT_EXISTS)
	}

	postService := post.NewPostService(postController.server.Redis)
	res := postService.GetPostBySlug(ctx.Request().Context(), key)

	return responses.Response(ctx, http.StatusOK, res)

}
