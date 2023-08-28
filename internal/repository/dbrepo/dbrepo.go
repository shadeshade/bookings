package dbrepo

import (
	"database/sql"
	"github.com/shadeshade/bookings-api/internal/config"
	"github.com/shadeshade/bookings-api/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(a *config.AppConfig, conn *sql.DB) repository.DatabaseRepo {
	return &postgresDBRepo{App: a, DB: conn}
}
