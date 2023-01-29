package main

import (
	"elprup/api/v1/router"
	"elprup/internal/app"
	"elprup/internal/config"
	"elprup/internal/logger"
	"elprup/repository"
	"net/http"
)

func main() {
	logger.Init()
	cfg := config.LoadConfig()
	repo := repository.NewPgRepository(&cfg.DbConfig)
	app := app.App{
		Repository: repo,
		Config:     cfg,
	}
	logger.Info(app)
	r := router.NewRouter()

	logger.Fatal(http.ListenAndServe(":3000", r))
}
