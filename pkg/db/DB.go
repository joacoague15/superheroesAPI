package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:root@localhost:5432/heroes_db"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return db
}
