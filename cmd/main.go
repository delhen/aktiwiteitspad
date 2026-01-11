package main

import (
	"github.com/gin-gonic/gin"
	"go_aktiwiteitspad/api/handlers"
)

func main() {
	handlers.InitializeDatabase()
	service := gin.Default()

	service.POST("/users", handlers.CreateUser)

	service.Run(":8080")
}
