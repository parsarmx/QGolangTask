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

func (postController *PostController) CreatePost(ctx echo.Context) error {
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
	fmt.Println(val)
	if val == 1 {
		return responses.ErrorResponse(ctx, http.StatusBadRequest, responses.SLUG_ALREADY_EXISTS)
	}

	postService := post.NewPostService(postController.server.Redis)
	done, err := postService.CreatePost(key, *req)
	if err != nil {
		fmt.Printf("(CreatePostController): %v \n", err)
		return responses.ErrorResponse(ctx, http.StatusBadGateway, responses.ERROR_DURNING_CREATE_POST)
	}

	response := responses.CreatePostResponse{
		Created: done,
		Status:  fmt.Sprintf("%v", http.StatusCreated),
	}

	return responses.Response(ctx, http.StatusCreated, response)

}
