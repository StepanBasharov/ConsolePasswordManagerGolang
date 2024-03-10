package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("core/database/data/database.db"), &gorm.Config{})

	if err != nil {
		panic("Connection failed")
	}
	db.AutoMigrate(&MasterPassword{}, &Passwords{})
	return db
}

func GetOrCreateMasterPassword() {
	db := CreateConnection()
	var mp []MasterPassword
	db.Find(&mp)
	if len(mp) == 0 {
		fmt.Println("Not Found Master Password")
	} else {
		fmt.Println("doup")
	}
}
