package services

import (
	"book-data-management-railway/models"
	"book-data-management-railway/repositories"
	"errors"
)

func GetAllBooks() ([]models.Book, error) {
	return repositories.GetAllBooks()
}

func GetBookByID(id int) (*models.Book, error) {
	return repositories.GetBookByID(id)
}

func GetBooksByCategoryID(id int) ([]models.Book, error) {
	return repositories.GetBooksByCategoryID(id)
}

func CreateBook(b models.Book, createdBy string) error {

	// VALIDASI
	if b.Title == "" {
		return errors.New("title is required")
	}

	if b.ReleaseYear < 1980 || b.ReleaseYear > 2024 {
		return errors.New("release year must be between 1980 and 2024")
	}

	if b.TotalPage <= 0 {
		return errors.New("total_page must be greater than 0")
	}

	// VALIDASI CATEGORY
	category, _ := repositories.GetCategoryByID(b.CategoryID)
	if category == nil {
		return errors.New("category not found")
	}

	// LOGIC THICKNESS
	if b.TotalPage > 100 {
		b.Thickness = "tebal"
	} else {
		b.Thickness = "tipis"
	}

	return repositories.CreateBook(b, createdBy)
}

func UpdateBook(id int, b models.Book, modifiedBy string) error {

	existing, err := repositories.GetBookByID(id)
	if err != nil || existing == nil {
		return errors.New("book not found")
	}

	// VALIDASI CATEGORY
	category, _ := repositories.GetCategoryByID(b.CategoryID)
	if category == nil {
		return errors.New("category not found")
	}

	if b.ReleaseYear < 1980 || b.ReleaseYear > 2024 {
		return errors.New("invalid release year")
	}

	if b.TotalPage > 100 {
		b.Thickness = "tebal"
	} else {
		b.Thickness = "tipis"
	}

	return repositories.UpdateBook(id, b, modifiedBy)
}

func DeleteBook(id int) error {
	book, err := repositories.GetBookByID(id)
	if err != nil || book == nil {
		return errors.New("book not found")
	}

	return repositories.DeleteBook(id)
}