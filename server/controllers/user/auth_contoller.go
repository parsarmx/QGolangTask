package user

import (
	"fmt"
	"net/http"
	"task/models"
	"task/responses"
	"task/selectors"
	s "task/server"
	"task/services/token"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	server *s.Server
}

func NewAuthController(server *s.Server) *AuthController {
	return &AuthController{server: server}
}

func (authController *AuthController) Login(ctx echo.Context) error {
	req := models.LoginRequest{}
	err := ctx.Bind(&req)
	if req.Username == "" || req.Password == "" {
		fmt.Println(err)
		return responses.ErrorResponse(
			ctx, http.StatusBadRequest, responses.INVALID_REQUEST_BODY,
		)
	}
	// check if username exists or not
	user := models.User{}
	userSelector := selectors.NewUserSelector(authController.server.DB)
	userSelector.GetUserByUsername(&user, req.Username)
	if user.ID == 0 || (bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil) {
		return responses.ErrorResponse(ctx, http.StatusBadRequest, responses.USER_DOES_NOT_EXISTS)
	}

	tokenService := token.NewTokenService(authController.server.Config)
	access, refresh, err := tokenService.GenerateToken(user.Username, user.ID)
	if err != nil {
		return responses.ErrorResponse(
			ctx, http.StatusBadGateway, responses.FAILED_TO_CREATE_AUTH_TOKENS,
		)
	}

	// generate success response
	response := responses.RegisterResponse{
		AcessToken:   access,
		RefreshToken: refresh,
		Status:       fmt.Sprintf("%v", http.StatusCreated),
	}

	return responses.Response(ctx, http.StatusOK, response)
}
