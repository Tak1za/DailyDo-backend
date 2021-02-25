package database

import "gorm.io/gorm"

// DB struct
type DB struct {
	DB *gorm.DB
}

// Initiate function initiates the database
func Initiate(d *gorm.DB) *DB {
	return &DB{
		DB: d,
	}
}
