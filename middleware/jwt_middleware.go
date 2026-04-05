package middleware

import (
	"book-data-management-railway/config"
	"book-data-management-railway/repositories"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var SECRET_KEY = []byte("secret")

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 1. ambil header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "missing token",
			})
			c.Abort()
			return
		}

		// 2. validasi format Bearer
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "invalid token format",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 3. parse JWT + validasi signing method
		parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return SECRET_KEY, nil
		})

		if err != nil || !parsedToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "invalid token",
			})
			c.Abort()
			return
		}

		// 4. ambil claims dengan aman
		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "invalid token claims",
			})
			c.Abort()
			return
		}

		// 5. ambil user_id dengan aman
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "invalid user_id in token",
			})
			c.Abort()
			return
		}

		userID := int(userIDFloat)

		// 6. cek token di DB (session validation)
		var exists bool
		query := `SELECT EXISTS (SELECT 1 FROM user_sessions WHERE token=$1)`
		err = config.DB.QueryRow(query, tokenString).Scan(&exists)

		if err != nil || !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "session not found",
			})
			c.Abort()
			return
		}

		// 7. ambil user dari DB
		user, err := repositories.GetUserByID(userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "user not found",
			})
			c.Abort()
			return
		}

		// 8. inject ke context
		c.Set("user", user)
		c.Set("user_id", userID)
		c.Set("token", tokenString)

		c.Next()
	}
}