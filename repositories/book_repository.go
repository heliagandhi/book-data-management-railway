package repositories

import (
	"book-data-management-railway/config"
	"book-data-management-railway/models"
)

func GetAllBooks() ([]models.Book, error) {
	rows, err := config.DB.Query(`
		SELECT 
			b.id, b.title, b.description, b.image_url,
			b.release_year, b.price, b.total_page, b.thickness,
			c.id, c.name
		FROM books b
		JOIN categories c ON b.category_id = c.id
		ORDER BY b.id ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var b models.Book
		var c models.Category

		err := rows.Scan(
			&b.ID,
			&b.Title,
			&b.Description,
			&b.ImageURL,
			&b.ReleaseYear,
			&b.Price,
			&b.TotalPage,
			&b.Thickness,
			&c.ID,
			&c.Name,
		)
		if err != nil {
			return nil, err
		}

		b.Category = c
		books = append(books, b)
	}

	return books, nil
}

func GetBookByID(id int) (*models.Book, error) {
	var b models.Book
	var c models.Category

	err := config.DB.QueryRow(`
		SELECT 
			b.id, b.title, b.description, b.image_url,
			b.release_year, b.price, b.total_page, b.thickness,
			c.id, c.name
		FROM books b
		JOIN categories c ON b.category_id = c.id
		WHERE b.id = $1
	`, id).Scan(
		&b.ID,
		&b.Title,
		&b.Description,
		&b.ImageURL,
		&b.ReleaseYear,
		&b.Price,
		&b.TotalPage,
		&b.Thickness,
		&c.ID,
		&c.Name,
	)

	if err != nil {
		return nil, err
	}

	b.Category = c
	return &b, nil
}

func GetBooksByCategoryID(categoryID int) ([]models.Book, error) {
	rows, err := config.DB.Query(`
		SELECT 
			b.id, b.title, b.description, b.image_url,
			b.release_year, b.price, b.total_page, b.thickness,
			c.id, c.name
		FROM books b
		JOIN categories c ON b.category_id = c.id
		WHERE c.id = $1
	`, categoryID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var b models.Book
		var c models.Category

		err := rows.Scan(
			&b.ID,
			&b.Title,
			&b.Description,
			&b.ImageURL,
			&b.ReleaseYear,
			&b.Price,
			&b.TotalPage,
			&b.Thickness,
			&c.ID,
			&c.Name,
		)
		if err != nil {
			return nil, err
		}

		b.Category = c
		books = append(books, b)
	}

	return books, nil
}

func CreateBook(b models.Book, createdBy string) error {
	_, err := config.DB.Exec(`
		INSERT INTO books 
		(title, category_id, description, image_url, release_year, price, total_page, thickness, created_by)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	`,
		b.Title,
		b.CategoryID,
		b.Description,
		b.ImageURL,
		b.ReleaseYear,
		b.Price,
		b.TotalPage,
		b.Thickness,
		createdBy,
	)

	return err
}

func UpdateBook(id int, b models.Book, modifiedBy string) error {
	_, err := config.DB.Exec(`
		UPDATE books SET
		title=$1,
		category_id=$2,
		description=$3,
		image_url=$4,
		release_year=$5,
		price=$6,
		total_page=$7,
		thickness=$8,
		modified_by=$9,
		modified_at=NOW()
		WHERE id=$10
	`,
		b.Title,
		b.CategoryID,
		b.Description,
		b.ImageURL,
		b.ReleaseYear,
		b.Price,
		b.TotalPage,
		b.Thickness,
		modifiedBy,
		id,
	)

	return err
}

func DeleteBook(id int) error {
	res, err := config.DB.Exec(`DELETE FROM books WHERE id=$1`, id)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return nil
	}

	return nil
}