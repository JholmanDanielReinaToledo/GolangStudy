package main

import (
	"fmt"
	"gorm/controllers"
	"gorm/db"

	"github.com/gin-gonic/gin"
)

func main() {
	Port := 8000
	router := gin.Default()

	db.ConnectDatabase()

	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts", controllers.FindPosts)
	router.GET("/posts/:id", controllers.FindPost)
	router.PATCH("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)

	fmt.Printf("Server Running in port %d \n", Port)
	router.Run("localhost:8080")
}
