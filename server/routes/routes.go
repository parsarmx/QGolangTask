package routes

import (
	s "task/server"

	"github.com/labstack/echo/v4/middleware"
)

func ConfigureRoutes(server *s.Server) {
	server.Echo.Use(middleware.Logger())

	r := server.Echo.Group("")
	_ = r
}
