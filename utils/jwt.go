package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SECRET_KEY = []byte("secret")

func GenerateJWT(userID int) (string, time.Time, error) {
	expiredAt := time.Now().Add(1 * time.Hour)

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expiredAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expiredAt, nil
}