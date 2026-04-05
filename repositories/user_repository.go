package repositories

import (
	"book-data-management-railway/config"
	"book-data-management-railway/models"
	"errors"
	"time"
)

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	err := config.DB.QueryRow(
		`SELECT id, username, password FROM users WHERE username=$1`,
		username,
	).Scan(&user.ID, &user.Username, &user.Password)

	if err != nil {
		return nil, nil // user tidak ditemukan
	}

	return &user, nil
}

func CreateUserSession(userID int, token string, expiredAt time.Time) error {
	query := `
	INSERT INTO user_sessions (user_id, token, expired_at)
	VALUES ($1, $2, $3)
	`
	_, err := config.DB.Exec(query, userID, token, expiredAt)
	return err
}

func GetAllUsers() ([]models.User, error) {
	rows, err := config.DB.Query(`SELECT id, username FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Username); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func CreateUser(username, password string) error {
	_, err := config.DB.Exec(
		`INSERT INTO users (username, password) VALUES ($1,$2)`,
		username, password,
	)
	return err
}

func UpdateUser(id int, username string) error {
	res, err := config.DB.Exec(
		`UPDATE users SET username=$1 WHERE id=$2`,
		username, id,
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("no rows updated")
	}

	return nil
}

func UpdatePassword(id int, password string) error {
	_, err := config.DB.Exec(
		`UPDATE users SET password=$1 WHERE id=$2`,
		password, id,
	)
	return err
}

func GetUserByID(id int) (*models.User, error) {
	var u models.User

	err := config.DB.QueryRow(
		`SELECT id, username FROM users WHERE id=$1`,
		id,
	).Scan(&u.ID, &u.Username)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func DeleteUser(id int) error {
	res, err := config.DB.Exec(`DELETE FROM users WHERE id=$1`, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return nil // handle di service
	}

	return nil
}