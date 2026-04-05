package controllers

import (
	"book-data-management-railway/models"
	"book-data-management-railway/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateBook godoc
// @Summary Create book
// @Tags Books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.Book true "Book Data"
// @Success 200 {object} map[string]string
// @Router /books [post]
func CreateBook(c *gin.Context) {
	var req models.Book

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": "invalid request",
		})
		return
	}

	user := c.MustGet("user").(*models.User)

	err := services.CreateBook(req, user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "book created",
	})
}

// GetBooks godoc
// @Summary Get all books
// @Tags Books
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Book
// @Router /books [get]
func GetBooks(c *gin.Context) {
	data, err := services.GetAllBooks()
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

// GetBookDetail godoc
// @Summary Get book detail
// @Tags Books
// @Produce json
// @Security BearerAuth
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Router /books/{id} [get]
func GetBookDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := services.GetBookByID(id)
	if err != nil || data == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"message": "book not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}

// UpdateBook godoc
// @Summary Update book
// @Tags Books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Book ID"
// @Param request body models.Book true "Book Data"
// @Success 200 {object} map[string]string
// @Router /books/{id} [put]
func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req models.Book
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": "invalid request",
		})
		return
	}

	user := c.MustGet("user").(*models.User)

	err := services.UpdateBook(id, req, user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "book updated",
	})
}

// DeleteBook godoc
// @Summary Delete book
// @Tags Books
// @Produce json
// @Security BearerAuth
// @Param id path int true "Book ID"
// @Success 200 {object} map[string]string
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := services.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "book deleted",
	})
}