package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB global instance of database
var DB *gorm.DB

// InitDB initial db
func InitDB() {
	db, err := gorm.Open("sqlite3", "db.sql")

	if err != nil {
		log.Fatalf("coudn't open database: %s", err.Error())
	}

	DB = db.LogMode(true)
}

// GetDB return an instance of DB
func GetDB() *gorm.DB {
	return DB
}
