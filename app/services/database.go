package services

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDatabase() {
	if nil == DB {
		dir, _ := os.Getwd()
		connection := sqlite.Open(dir + "/sqlite.db")
		config := &gorm.Config{
			NamingStrategy: schema.NamingStrategy{SingularTable: true},
		}
		if db, err := gorm.Open(connection, config); nil == err {
			DB = db.Debug()
		}
	}
}
