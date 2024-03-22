package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task/models"
	"task/responses"
	"task/selectors"
	s "task/server"
	"task/services/token"
	"task/services/user"

	"github.com/labstack/echo/v4"
)

type RegisterHandler struct {
	server *s.Server
}

func NewRegisterHandler(server *s.Server) *RegisterHandler {
	return &RegisterHandler{server: server}
}

func (registerHandler *RegisterHandler) Register(ctx echo.Context) error {

	// binding body
	req := models.RegisterRequest{}
	err := ctx.Bind(&req)
	if req.Username == "" || req.Password == "" {
		return responses.ErrorResponse(
			ctx, http.StatusBadRequest, responses.INVALID_REQUEST_BODY,
		)
	}

	// check if username exists or not
	existsUser := models.User{}
	userSelector := selectors.NewUserSelector(registerHandler.server.DB)
	userSelector.GetUserByUsername(&existsUser, req.Username)
	if existsUser.ID != 0 {
		return responses.ErrorResponse(ctx, http.StatusBadRequest, responses.USER_ALREADY_EXISTS)
	}

	// trying to create user
	userService := user.NewUserService(registerHandler.server.DB)
	user, err := userService.RegisterUser(req)
	if err != nil {
		return responses.ErrorResponse(
			ctx, http.StatusBadRequest, responses.FAILED_TO_CREATE_USER,
		)
	}

	// generate token
	tokenService := token.NewTokenService(registerHandler.server.Config)
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

	data, err := json.Marshal(&response)
	if err != nil {
		return responses.ErrorResponse(
			ctx, http.StatusBadGateway, responses.FAILED_TO_MARSHAL,
		)
	}

	return responses.MessageResponse(ctx, http.StatusCreated, string(data))
}
