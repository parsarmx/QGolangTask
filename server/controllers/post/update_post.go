package post

import (
	"fmt"
	"net/http"
	"task/models"
	"task/responses"
	"task/selectors"
	"task/services/post"

	"github.com/labstack/echo/v4"
)

func (postController *PostController) UpdatePost(ctx echo.Context) error {
	req := new(models.PostRequest)
	slug := ctx.Param("slug")
	req.Slug = slug

	err := ctx.Bind(&req)
	if err != nil {
		return err
	}

	if err := req.Validate(); err != nil {
		fmt.Println(err)
		return responses.ErrorResponse(ctx, http.StatusBadRequest, responses.REQUIRED_FIELDS_ARE_EMPTY)
	}

	key := fmt.Sprintf("post:%v", req.Slug)
	postSelector := selectors.NewPostSelector(postController.server.Redis)
	val, _ := postSelector.IsSlugExists(ctx.Request().Context(), key)
	if val != 1 {
		return responses.ErrorResponse(ctx, http.StatusBadRequest, responses.SLUG_DOES_NOT_EXISTS)
	}

	postService := post.NewPostService(postController.server.Redis)
	done, err := postService.UpdatePost(ctx.Request().Context(), key, *req)
	if err != nil {
		fmt.Printf("(UpdatePostController): %v \n", err)
		return responses.ErrorResponse(ctx, http.StatusBadGateway, responses.ERROR_DURNING_UPDATE_POST)
	}

	response := responses.UpdatePostResponse{
		Updated: done,
		Status:  fmt.Sprintf("%v", http.StatusCreated),
	}

	return responses.Response(ctx, http.StatusCreated, response)
}
