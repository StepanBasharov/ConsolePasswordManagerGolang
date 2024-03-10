package main

import (
	"ConsolePasswordManager/core"
	"ConsolePasswordManager/core/database"
)

func main() {
	result := core.CreateHashMasterPassword("test_password")
	testValid := core.CheckMasterPasswordValid("test_password", result)
	println(testValid)
	database.GetOrCreateMasterPassword()
}
