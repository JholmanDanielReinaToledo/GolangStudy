package reports

import (
	"gorm/db"
	"gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateXlsx(c *gin.Context) {
	var posts []models.Post
	db.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{"data": posts})
}
