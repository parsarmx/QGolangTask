package routes

import (
	s "task/server"
	"task/server/controllers/user"

	"github.com/labstack/echo/v4/middleware"
)

func ConfigureRoutes(server *s.Server) {
	registerHandler := user.NewRegisterHandler(server)
	authHandler := user.NewAuthHandler(server)

	server.Echo.Use(middleware.Logger())

	server.Echo.POST("/register", registerHandler.Register)
	server.Echo.POST("/login", authHandler.Login)
}
