package initializers

import "github.com/Abdullahxz/gin-jwt/models"

func MigrateDB() {
	DB.AutoMigrate(&models.User{})
}
