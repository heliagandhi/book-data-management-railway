package repositories

import (
	"book-data-management-railway/config"
	"book-data-management-railway/models"
)

func GetAllCategories() ([]models.Category, error) {
	rows, err := config.DB.Query(`SELECT id, name FROM categories`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category

	for rows.Next() {
		var c models.Category
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}

func GetCategoryByID(id int) (*models.Category, error) {
	var c models.Category

	err := config.DB.QueryRow(
		`SELECT id, name FROM categories WHERE id=$1`,
		id,
	).Scan(&c.ID, &c.Name)

	if err != nil {
		return nil, nil // ⬅️ ini penting
	}

	return &c, nil
}

func CreateCategory(name, createdBy string) error {
	_, err := config.DB.Exec(
		`INSERT INTO categories (name, created_by) VALUES ($1, $2)`,
		name,
		createdBy,
	)
	return err
}

func UpdateCategory(id int, name, modifiedBy string) error {
	_, err := config.DB.Exec(
		`UPDATE categories 
		 SET name=$1, modified_by=$2, modified_at=NOW()
		 WHERE id=$3`,
		name, modifiedBy, id,
	)
	return err
}

func DeleteCategory(id int) error {
	res, err := config.DB.Exec(`DELETE FROM categories WHERE id=$1`, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return nil // kita handle di service
	}

	return nil
}