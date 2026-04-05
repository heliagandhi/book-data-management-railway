package services

import (
	"book-data-management-railway/models"
	"book-data-management-railway/repositories"
	"book-data-management-railway/utils"
	"errors"
)

func GetAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}

func GetUserByID(id int) (*models.User, error) {
	return repositories.GetUserByID(id)
}

func CreateUser(username, password string) error {
	if username == "" || password == "" {
		return errors.New("username and password required")
	}

	existing, _ := repositories.GetUserByUsername(username)
	if existing != nil {
		return errors.New("username already exists")
	}

	hashed, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	return repositories.CreateUser(username, hashed)
}

func UpdateUser(id int, username string) error {
	user, err := repositories.GetUserByID(id)
	if err != nil || user == nil {
		return errors.New("user not found")
	}

	if username == "" {
		return errors.New("username required")
	}

	// optional: cek duplicate username saat update
	existing, _ := repositories.GetUserByUsername(username)
	if existing != nil && existing.ID != id {
		return errors.New("username already used")
	}

	return repositories.UpdateUser(id, username)
}

func UpdatePassword(id int, password string) error {
	if password == "" {
		return errors.New("password required")
	}

	user, err := repositories.GetUserByID(id)
	if err != nil || user == nil {
		return errors.New("user not found")
	}

	hashed, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	return repositories.UpdatePassword(id, hashed)
}

func DeleteUser(id int) error {
	user, err := repositories.GetUserByID(id)
	if err != nil || user == nil {
		return errors.New("user not found")
	}

	return repositories.DeleteUser(id)
}