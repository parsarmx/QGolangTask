package routes

import (
	s "task/server"
	"task/server/controllers/user"

	"github.com/labstack/echo/v4/middleware"
)

func ConfigureRoutes(server *s.Server) {
	authControler := user.NewRegisterHandler(server)

	server.Echo.Use(middleware.Logger())

	server.Echo.POST("/register", authControler.Register)
}
