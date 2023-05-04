package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	passwordBytes := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hashString := string(hash)
	return hashString, nil
}

func ComparePasswords(hashedPassword string, password string) error {
	hashedPasswordBytes := []byte(hashedPassword)

	passwordBytes := []byte(password)

	err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, passwordBytes)
	if err != nil {
		return err
	}

	return nil
}
