package main

import (
	"elprup/api/v1/router"
	"elprup/internal/config"
	"elprup/internal/logger"
	"net/http"
)

func main() {
	logger.Init()
	cfg := config.LoadConfig()
	logger.Info(cfg)
	r := router.NewRouter()

	logger.Fatal(http.ListenAndServe(":3000", r))
}
