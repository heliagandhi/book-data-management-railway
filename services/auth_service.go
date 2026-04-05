package services

import (
	"book-data-management-railway/repositories"
	"book-data-management-railway/utils"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Login(username, password string) (int, string, time.Time, error) {
	user, err := repositories.GetUserByUsername(username)
	if err != nil {
		return 0, "", time.Time{}, errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return 0, "", time.Time{}, errors.New("invalid username or password")
	}

	token, expiredAt, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return 0, "", time.Time{}, err
	}

	err = repositories.CreateUserSession(user.ID, token, expiredAt)
	if err != nil {
		return 0, "", time.Time{}, err
	}

	return user.ID, token, expiredAt, nil
}