package routes

import (
	"task/middlewares"
	s "task/server"
	"task/server/controllers/post"
	"task/server/controllers/user"

	"github.com/labstack/echo/v4/middleware"
)

func ConfigureRoutes(server *s.Server) {
	registerHandler := user.NewRegisterController(server)
	authHandler := user.NewAuthController(server)
	postHandler := post.NewPostController(server)

	server.Echo.Use(middleware.Logger())

	server.Echo.POST("/register/", registerHandler.Register)
	server.Echo.POST("/login/", authHandler.Login)

	// posts
	server.Echo.POST("/post/:slug/create/", postHandler.CreatePost, middlewares.AuthMiddleware)
	server.Echo.PUT("/post/:slug/update/", postHandler.UpdatePost, middlewares.AuthMiddleware)
	server.Echo.GET("/post/:slug/", postHandler.GetPostBySlug, middlewares.AuthMiddleware)
	server.Echo.GET("/post/", postHandler.GetAllPosts, middlewares.AuthMiddleware)
}
