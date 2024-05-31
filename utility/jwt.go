package utility

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const key = "FnNRSw7ltG"

func GenerateJwtToken(emailId string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"email":  emailId,
		"userId": userId,
		"expiry": time.Now().Add(time.Hour * 2)})

	return token.SignedString([]byte(key))
}

func VerifyToken(token string) (int, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected Signing method")
		}
		return []byte(key), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse token")
	}

	if !parsedToken.Valid {
		return 0, errors.New("Invalid token")
	}

	data, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid token")
	}
	userIdFloat, ok := data["userId"].(float64)
	return int(userIdFloat), nil

}
