package main

import (
	"gorm/controllers"
	"gorm/db"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db.ConnectDatabase()

	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts", controllers.FindPosts)
	router.GET("/posts/:id", controllers.FindPost)
	router.PATCH("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)

	router.Run("localhost:8080")
}
