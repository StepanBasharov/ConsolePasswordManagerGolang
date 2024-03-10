package cipher_manager

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
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

func EncryptUserPassword(userPassword string, masterPassword string) string {
	masterPassword = base64.StdEncoding.EncodeToString([]byte(masterPassword))
	userPasswordBytes := []byte(userPassword)
	masterPasswordBytes := []byte(masterPassword)

	block, err := aes.NewCipher(masterPasswordBytes)
	if err != nil {
		fmt.Printf("Оно 1 %v", err)
		panic("Enc error")
	}

	cipherText := make([]byte, aes.BlockSize+len(userPasswordBytes))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Printf("Оно 2 %v", err)
		panic("Enc error")
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], userPasswordBytes)

	return base64.StdEncoding.EncodeToString(cipherText)
}

func DecryptUserPassword(masterPassword string, userPassword string) string {
	masterPassword = base64.StdEncoding.EncodeToString([]byte(masterPassword))
	masterPasswordBytes := []byte(masterPassword)
	cipherText, err := base64.StdEncoding.DecodeString(userPassword)
	if err != nil {
		panic("could not base64 decode")
	}

	block, err := aes.NewCipher(masterPasswordBytes)
	if err != nil {
		panic("could not create new cipher")
	}

	if len(cipherText) < aes.BlockSize {
		panic("invalid ciphertext block size")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText)
}
