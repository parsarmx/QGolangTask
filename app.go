package application

import (
	"log"
	"task/configs"
	"task/server"
	"task/server/routes"
)

func Start(cfg *configs.Config) {

	app := server.NewServer(cfg)

	routes.ConfigureRoutes(app)

	err := app.Start(cfg.HTTP.Port)

	if err != nil {
		log.Fatal("Port already used")
	}
}
