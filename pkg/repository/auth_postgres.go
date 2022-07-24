package repository

import (
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db sqlx.DB
}

func newAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
