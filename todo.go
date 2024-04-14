package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
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
	dsn := "root@tcp(localhost:3306)/tododb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	err = db.AutoMigrate(&TodoItem{})
	if err != nil {
		log.Fatal("Failed to migrate the schema")
	}
}
