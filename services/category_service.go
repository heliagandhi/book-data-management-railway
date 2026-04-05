package services

import (
	"book-data-management-railway/models"
	"book-data-management-railway/repositories"
	"errors"
)

func GetAllCategories() ([]models.Category, error) {
	return repositories.GetAllCategories()
}

func GetCategoryByID(id int) (*models.Category, error) {
	return repositories.GetCategoryByID(id)
}

func CreateCategory(name, createdBy string) error {
	if name == "" {
		return errors.New("category name is required")
	}

	return repositories.CreateCategory(name, createdBy)
}

func UpdateCategory(id int, name, modifiedBy string) error {
	cat, err := repositories.GetCategoryByID(id)
	if err != nil || cat == nil {
		return errors.New("category not found")
	}

	if name == "" {
		return errors.New("category name is required")
	}

	return repositories.UpdateCategory(id, name, modifiedBy)
}

func DeleteCategory(id int) error {
	cat, err := repositories.GetCategoryByID(id)
	if err != nil || cat == nil {
		return errors.New("category not found")
	}

	// books, _ := repositories.GetBooksByCategoryID(id)
	// if len(books) > 0 {
	// 	return errors.New("cannot delete category with existing books")
	// }

	return repositories.DeleteCategory(id)
}