package repository

import (
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type DbContext struct {
}

func NewDbContext() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=1905 dbname=rentacarorg port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return db, err
	}
	return db, nil
}
