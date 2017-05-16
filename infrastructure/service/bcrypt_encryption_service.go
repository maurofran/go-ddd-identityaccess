package service

import "golang.org/x/crypto/bcrypt"

type BCryptEncryptionService struct {}

func (es *BCryptEncryptionService) EncryptValue(plainText string) (string, error) {
	bytes := []byte(plainText)
	encoded, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
	return string(encoded), err
}