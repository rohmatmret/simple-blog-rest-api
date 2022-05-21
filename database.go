package main

import (
	"log"

	"github.com/simple-blog/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./post_db.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Migration(db)
	return db
}

func Migration(db *gorm.DB) {
	// Migrate the schema
	err := db.AutoMigrate(domain.Post{})
	if err != nil {
		log.Fatal("error when migrate", err)
	}
}
