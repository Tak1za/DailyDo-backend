package database

import (
	"database/sql"
)

// Store interface
type Store interface {
	Get(ID int) (string, error)
}

type store struct {
	db *sql.DB
}

func (d *store) Get(ID int) (string, error) {
	return "task1", nil
}
