package controllers

import (
	"github.com/MatheusMeloAntiquera/api-go/src/database"
	"github.com/MatheusMeloAntiquera/api-go/src/models"
	"github.com/gin-gonic/gin"
)

func UserIndex(context *gin.Context) {
	var users []models.User
	database.DbCon.Find(&users)
	context.JSON(200, gin.H{
		"success": true,
		"data":    users,
	})
}

func UserShow(context *gin.Context) {
	var user models.User
	database.DbCon.First(&user, context.Param("id"))
	if checkUserExist(user, context) {
		context.JSON(200, gin.H{
			"success": true,
			"data":    user,
		})
	}
}

func UserCreate(context *gin.Context) {
	user := models.User{Name: context.PostForm("name")}
	database.DbCon.Create(&user)
	context.JSON(200, gin.H{
		"success": true,
		"data":    user,
	})
}

func UserUpdate(context *gin.Context) {
	var user models.User
	database.DbCon.First(&user, context.Param("id"))
	if checkUserExist(user, context) {
		user.Name = context.PostForm("name")
		database.DbCon.Save(&user)
		context.JSON(200, gin.H{
			"success": true,
			"data":    user,
		})
	}
}

func UserDelete(context *gin.Context) {
	var user models.User
	database.DbCon.First(&user, context.Param("id"))

	if checkUserExist(user, context) {
		database.DbCon.Delete(&user)
		context.JSON(200, gin.H{
			"success": true,
			"message": "User deleted successfully",
		})
	}
}

func checkUserExist(user models.User, context *gin.Context) bool {
	if (user == models.User{}) {
		context.JSON(200, gin.H{
			"success": false,
			"message": "User not found",
		})

		return false
	}
	return true
}
