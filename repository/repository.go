package repository

import (
	"elprup/internal/config"
	"elprup/internal/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository interface {
}

type PgRepository struct {
	DB *gorm.DB
}

func NewPgRepository(dbConfig *config.DbConfig) *PgRepository {
	db, err := gorm.Open(postgres.Open(dbConfig.Dsn))
	if err != nil {
		logger.Panic(err)
	}
	return &PgRepository{
		DB: db,
	}
}
