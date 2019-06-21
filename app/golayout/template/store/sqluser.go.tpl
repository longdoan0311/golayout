package store

import (
	"github.com/jmoiron/sqlx"

	"{{.ModName}}"
)

type sqlUser struct {
	db *sqlx.DB
}

func (s *sqlUser) GetAll() ([]{{.ProjName}}.User, error) {
	return make([]{{.ProjName}}.User, 0), nil
}