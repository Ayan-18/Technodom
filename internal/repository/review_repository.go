package repository

import (
	"database/sql"
	"log/slog"
)

type ReviewRepository struct {
	db     *sql.DB
	logger *slog.Logger
}
