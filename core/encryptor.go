package core

import (
	"crypto/sha256"
	"encoding/hex"
)

func CreateHashMasterPassword(password string) string {
	key := sha256.New()
	key.Write([]byte(password))
	keyBs := key.Sum(nil)
	keyBsString := hex.EncodeToString(keyBs)
	return keyBsString
}

func CheckMasterPasswordValid(password string, sha256Password string) bool {
	passwordHash := CreateHashMasterPassword(password)
	if passwordHash != sha256Password {
		return false
	}
	return true
}
