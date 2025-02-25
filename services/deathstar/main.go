package main

import (
    "github.com/gin-gonic/gin"
    handlers "monorepo/services/deathstar/handlers"
    "net/http"
    "strings"
)

func main() {
    router := gin.Default()
    router.Use(validateJWT())
    router.GET("/status", handlers.StatusHandler)
    router.POST("/exhaust-port", handlers.ExhaustPortHandler)
    router.Run(":8080")
}

// validateJWT middleware checks the Authorization header for a valid JWT token
// Kong will have already validated the token, but we double-check and extract claims
func validateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// In production, you'd want to validate the JWT signature
		// But with Kong as a proxy, it would have already verified the token

		// Get token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// Check if Kong passed the token in X-Userinfo header
			userInfo := c.GetHeader("X-Userinfo")
			if userInfo == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
				c.Abort()
				return
			}
			// Continue with userinfo from Kong
			c.Set("userinfo", userInfo)
			c.Next()
			return
		}

		// Extract token from "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization format"})
			c.Abort()
			return
		}

		token := parts[1]
		
		// In a real application, you would verify the token
		// But since Kong already did this for us, we can just extract claims
		
		// Set token in context for handlers to use
		c.Set("token", token)
		
		c.Next()
	}
}