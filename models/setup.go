package models

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&Book{})

	DB = database
}
