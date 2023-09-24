package main

import (
	"github.com/Abdullahxz/gin-jwt/controllers"
	"github.com/Abdullahxz/gin-jwt/initializers"
	"github.com/Abdullahxz/gin-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.MigrateDB()
}
func main() {
	r := gin.Default()

	// Health check endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.LogIn)

	// Demonstrates jwt authentication using middleware
	r.GET("/dummy", middleware.Auth, controllers.DummyController)

	r.Run()
}
