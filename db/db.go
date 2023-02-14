package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func InitDB() (*DB, error) {
	db, err := gorm.Open(postgres.Open("host=localhost user=pg password=password dbname=pg port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Person{})

	return &DB{
		db,
	}, nil
}
