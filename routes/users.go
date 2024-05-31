package routes

import (
	"events-management/models"
	"events-management/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}
	err = user.SaveUser()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User saved successfully", "user": user})
}

func loginUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	id, err := user.ValidateUser()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authorize user"})
		return
	}

	token, err := utility.GenerateJwtToken(user.Email, id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate user token"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User validated successfully", "token": token})
}
