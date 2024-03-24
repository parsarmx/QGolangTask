package post

import s "task/server"

type PostController struct {
	server *s.Server
}

func NewPostController(server *s.Server) *PostController {
	return &PostController{server: server}
}
