package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
}

func main() {
	router := gin.Default()

	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/api-go?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	router.GET("/user", func(context *gin.Context) {
		var users []User
		db.Find(&users)
		context.JSON(200, gin.H{
			"success": true,
			"data":    users,
		})
	})

	router.GET("/user/:id", func(context *gin.Context) {
		var user User
		db.First(&user, context.Param("id"))
		context.JSON(200, gin.H{
			"success": true,
			"data":    user,
		})
	})

	router.POST("/user/", func(context *gin.Context) {
		user := User{Name: context.PostForm("name")}
		db.Create(&user)
		context.JSON(200, gin.H{
			"success": true,
			"data":    user,
		})
	})

	router.PUT("/user/:id", func(context *gin.Context) {
		var user User
		db.First(&user, context.Param("id"))
		user.Name = context.PostForm("name")
		db.Save(&user)
		context.JSON(200, gin.H{
			"success": true,
			"data":    user,
		})
	})

	router.DELETE("/user/:id", func(context *gin.Context) {
		var user User
		db.First(&user, context.Param("id"))
		db.Delete(&user)
		context.JSON(200, gin.H{
			"success": true,
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
