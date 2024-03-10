package manager

import (
	"ConsolePasswordManager/command_line_manager"
	"ConsolePasswordManager/core/database"
	"fmt"
)

func UserLogin() (bool, string) {
	for i := 0; i <= 3; i++ {
		isValid, password := database.GetOrCreateMasterPassword()
		if isValid == true {
			return true, password
		} else {
			fmt.Println("Invalid Master Password, try again")
		}
	}
	return false, ""
}

func Entry() {
	command_line_manager.UserConsoleHello()
	isValid, password := UserLogin()
	if isValid == true {
		if len(password) > 0 {
			command_line_manager.ReadOrCreatePassword(password)
		} else {
			fmt.Println("New Password Created")
		}
	} else {
		fmt.Println("Overshoot, bye")
	}
}
