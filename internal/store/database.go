package store

import "gorm.io/gorm"
import "gorm.io/driver/sqlite"

var DB *gorm.DB

func Init(dsn string){
	var err error
	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&User{})
}