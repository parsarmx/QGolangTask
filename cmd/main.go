package main

import (
	application "task"
	"task/configs"
	"task/migrations"
)

func main() {

	cfg := configs.NewConfig()

	migrations.Migrate(&cfg.DB)

	application.Start(cfg)

}
