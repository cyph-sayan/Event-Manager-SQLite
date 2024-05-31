package utility

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(bytes), err
}

func ValidatePassword(password, retrievedPassWord string) error {
	err := bcrypt.CompareHashAndPassword([]byte(retrievedPassWord), []byte(password))
	return err
}
