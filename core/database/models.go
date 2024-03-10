package database

import "gorm.io/gorm"

type MasterPassword struct {
	gorm.Model
	Password string
}

type Passwords struct {
	gorm.Model
	Source       string
	PasswordHash string
}
