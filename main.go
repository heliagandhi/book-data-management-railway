// @title Book Data Management API
// @version 1.0
// @description API for managing books and categories
// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"book-data-management-railway/config"
	"book-data-management-railway/controllers"
	"book-data-management-railway/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "book-data-management-railway/docs"

	"github.com/gin-gonic/gin"
)

func main() {
	// init DB
	config.ConnectDB()

	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/users/login", controllers.Login)

		api.GET("/users", controllers.GetUsers)
		api.POST("/users", controllers.CreateUser)
		api.GET("/users/:id", controllers.GetUserDetail)
		api.PUT("/users/:id", controllers.UpdateUser)
		api.PUT("/users/:id/password", controllers.UpdatePassword)
		api.DELETE("/users/:id", controllers.DeleteUser)

		// protected routes
		protected := api.Group("/")
		protected.Use(middleware.JWTMiddleware())
		{
			protected.GET("/test", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "authorized"})
			})

			protected.GET("/categories", controllers.GetCategories)
			protected.POST("/categories", controllers.CreateCategory)
			protected.PUT("/categories/:id", controllers.UpdateCategory)
			protected.GET("/categories/:id", controllers.GetCategoryDetail)
			protected.GET("/categories/:id/books", controllers.GetBooksByCategory)
			protected.DELETE("/categories/:id", controllers.DeleteCategory)

			protected.GET("/books", controllers.GetBooks)
			protected.POST("/books", controllers.CreateBook)
			protected.PUT("/books/:id", controllers.UpdateBook)
			protected.GET("/books/:id", controllers.GetBookDetail)
			protected.DELETE("/books/:id", controllers.DeleteBook)
			
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}