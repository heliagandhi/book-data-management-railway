package controllers

import (
	"book-data-management-railway/models"
	"book-data-management-railway/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateCategory godoc
// @Summary Create category
// @Tags Categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]string true "Category Name"
// @Success 200 {object} map[string]string
// @Router /categories [post]
func CreateCategory(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
	}

	// 1. validasi request dulu
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": "invalid request",
		})
		return
	}
	// 2. ambil user dari middleware
	user := c.MustGet("user").(*models.User)
	// 3. kirim ke service
	err := services.CreateCategory(req.Name, user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "category created",
	})
}

// GetCategories godoc
// @Summary Get all categories
// @Tags Categories
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Category
// @Router /categories [get]
func GetCategories(c *gin.Context) {
	data, err := services.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}

// GetCategoryDetail godoc
// @Summary Get category detail
// @Tags Categories
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Success 200 {object} models.Category
// @Router /categories/{id} [get]
func GetCategoryDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := services.GetCategoryByID(id)
	if err != nil || data == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"message": "category not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}

// GetBooksByCategory godoc
// @Summary Get books by category
// @Tags Categories
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Success 200 {array} models.Book
// @Router /categories/{id}/books [get]
func GetBooksByCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := services.GetBooksByCategoryID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"message": err.Error(),
		})
		return
	}

	if len(data) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"message": "no books found for this category",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}

// UpdateCategory godoc
// @Summary Update category
// @Tags Categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Param request body map[string]string true "Category Name"
// @Success 200 {object} map[string]string
// @Router /categories/{id} [put]
func UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": "invalid id",
		})
		return
	}

	var req struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": "invalid request",
		})
		return
	}

	// ambil user
	user := c.MustGet("user").(*models.User)

	err = services.UpdateCategory(id, req.Name, user.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "category updated",
	})
}

// DeleteCategory godoc
// @Summary Delete category
// @Tags Categories
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]string
// @Router /categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := services.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "category deleted",
	})
}