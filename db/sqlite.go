package db

import "gorm.io/gorm"
import "gorm.io/driver/sqlite"

type DB struct {
	*gorm.DB
}

func New() (DB, error) {
	db, err := gorm.Open(sqlite.Open("logger.db"), &gorm.Config{})
	return DB{db}, err
}

