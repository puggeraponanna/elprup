package app

import (
	"elprup/internal/config"
	"elprup/repository"
)

type App struct {
	Repository repository.Repository
	Config     *config.Config
}
