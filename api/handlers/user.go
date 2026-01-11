package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go_aktiwiteitspad/api/models"
)

func CreateUser(context *gin.Context) {
	var user models.User

	// Bind request body
	if err := context.ShouldBindJSON(&user); err != nil {
		models.ResponseJSON(context, http.StatusBadRequest, "Invalid input", nil)

		return
	}

	user.Create(DB)

	models.ResponseJSON(context, http.StatusCreated, "User created successfully", user)
}