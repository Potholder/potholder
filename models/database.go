package models

import (
	"sync"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var once sync.Once

func getDB() *gorm.DB {
	once.Do(func() {
		dbInstance, err := gorm.Open("sqlite3", "potholder.sqlite3")
		if err != nil {
			panic("Failed to connect database")
		}
		db = dbInstance
	})
	return db
}
