package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type TodoItem struct {
	gorm.Model
	Title       string
	Description string
	completed   bool
}

var db *gorm.DB

func initDB() {
	var err error
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable is not set.")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = db.AutoMigrate(&TodoItem{})
	if err != nil {
		log.Fatal("Failed to migrate the schema")
	}
}
