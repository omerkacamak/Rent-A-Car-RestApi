package repository

import (
	"log"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDbContext() (*gorm.DB, error) {
	var err error
	dsn := "host=localhost user=postgres password=1905 dbname=rentacarorg port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	println("--------->::: eskicontext ")
	if err != nil {
		return DB, err
	}
	return DB, nil
}

func ConnectToDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=1905 dbname=rentacarorg port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	println("--------->::: yenicontext ")
	if err != nil {
		log.Fatal(err.Error())
	}

}
