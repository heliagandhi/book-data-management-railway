package controllers

import (
	"book-data-management-railway/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Login user to get token
// @Tags Users
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login Request"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Router /users/login [post]
func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "invalid request",
		})
		return
	}

	_, token, expiredAt, err := services.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "200",
		"message":     "success",
		"accessToken": token,
		"expiredAt":   expiredAt,
	})
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CreateUser godoc
// @Summary Register user
// @Tags Users
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Register Request"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": "invalid request",
		})
		return
	}

	err := services.CreateUser(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// GetUsers godoc
// @Summary Get all users
// @Tags Users
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func GetUsers(c *gin.Context) {
	data, err := services.GetAllUsers()
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

// GetUserDetail godoc
// @Summary Get user detail
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func GetUserDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := services.GetUserByID(id)
	if err != nil || data == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"message": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}

// UpdateUser godoc
// @Summary Update user username
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body LoginRequest true "Update Request"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": "invalid id",
		})
		return
	}

	var req struct {
		Username string `json:"username"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": "invalid request",
		})
		return
	}

	err = services.UpdateUser(id, req.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "user updated",
	})
}

// UpdatePassword godoc
// @Summary Update user password
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body LoginRequest true "Password Request"
// @Success 200 {object} map[string]string
// @Router /users/{id}/password [put]
func UpdatePassword(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req struct {
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": "invalid request",
		})
		return
	}

	err := services.UpdatePassword(id, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "password updated",
	})
}

// DeleteUser godoc
// @Summary Delete user
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": "invalid id",
		})
		return
	}

	err = services.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "user deleted",
	})
}