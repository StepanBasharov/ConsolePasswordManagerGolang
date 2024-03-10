package command_line_manager

import (
	"ConsolePasswordManager/cipher_manager"
	"ConsolePasswordManager/core/database"
	"fmt"
)

func UserConsoleHello() {
	fmt.Printf("Console Password Manager v0.0.1\n")
}

func ReadOrCreatePassword(masterPassword string) {
	var option string
	fmt.Println("Please, select option:\n1. Show my passwords\n2. Create new Password")
	fmt.Print("Select option: ")
	fmt.Scan(&option)
	db := database.CreateConnection()
	if option == "1" {
		var passwords []database.Passwords
		db.Find(&passwords)
		for index, item := range passwords {
			fmt.Printf("%d - Sourece: %s - Password: %s\n", index, item.Source, cipher_manager.DecryptUserPassword(masterPassword, item.PasswordHash))
		}
	} else if option == "2" {
		var source string
		var password string
		fmt.Print("Enter a source for password: ")
		fmt.Scan(&source)
		fmt.Print("Enter a password: ")
		fmt.Scan(&password)
		encPassword := cipher_manager.EncryptUserPassword(password, masterPassword)
		db.Create(&database.Passwords{Source: source, PasswordHash: encPassword})

	} else {
		fmt.Println("Bad Command")
	}
}
