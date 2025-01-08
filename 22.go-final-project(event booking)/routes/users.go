package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func signUp(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse the request body."})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't create the user."})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User Created Successfully", "data": user})
}

func login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse the request body."})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Logged In successfully"})
}
