package controllers

import (
	"gorm/db"
	"gorm/models"
	"gorm/reports"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateXlsxPost(c *gin.Context) {
	var posts []models.Post
	db.DB.Find(&posts)

	nameSheet := "Posts"
	headers := []string{"Id", "Title", "Content", "CreatedAt"}

	reports.GenerateXlsxGeneric(nameSheet, posts, headers)

	c.JSON(http.StatusOK, gin.H{"url": ""})
}
