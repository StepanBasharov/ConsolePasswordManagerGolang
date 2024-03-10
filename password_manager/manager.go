package password_manager

import (
	"ConsolePasswordManager/cipher_manager"
	"ConsolePasswordManager/core/database"
	"fmt"
)

func GetPasswords(masterPassword string) {
	var passwords []database.Passwords
	db := database.CreateConnection()
	db.Find(&passwords)
	for index, item := range passwords {
		fmt.Printf("%s - Sourece: %s - Password: %s", index, item.Source, cipher_manager.DecryptUserPassword(masterPassword, item.PasswordHash))
	}
}
