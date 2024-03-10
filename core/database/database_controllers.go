package database

import (
	"ConsolePasswordManager/cipher_manager"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetUserMasterPassword() string {
	var MasterPassword string
	fmt.Print("Please input Master Password: ")
	fmt.Scan(&MasterPassword)
	return MasterPassword
}

func GetNewMasterPassword() string {
	var newMasterPassword string
	fmt.Print("Please input new Master Password: ")
	fmt.Scan(&newMasterPassword)
	if len(newMasterPassword) < 12 || len(newMasterPassword) > 12 {
		fmt.Println("Mater Password must be at least 12 characters")
		return GetNewMasterPassword()
	}
	return newMasterPassword
}

func CreateConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("core/database/data/database.db"), &gorm.Config{})

	if err != nil {
		panic("Connection failed")
	}
	db.AutoMigrate(&MasterPassword{}, &Passwords{})
	return db
}

func GetOrCreateMasterPassword() (bool, string) {
	db := CreateConnection()
	var mp []MasterPassword
	db.Find(&mp)
	if len(mp) == 0 {
		fmt.Println("Not Found Master Password")
		newMasterPassword := GetNewMasterPassword()
		newMasterPasswordHash := cipher_manager.CreateHashMasterPassword(newMasterPassword)
		db.Create(&MasterPassword{Password: newMasterPasswordHash})
	} else {
		password := GetUserMasterPassword()
		isCorrectMasterPassword := cipher_manager.CheckMasterPasswordValid(password, mp[0].Password)
		if isCorrectMasterPassword {
			return true, password
		} else {
			return false, ""
		}
	}
	return true, ""
}
