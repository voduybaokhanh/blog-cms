package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/voduybaokhanh/blog-cms/models"
)

func GetTags(c *gin.Context) {
	var tags []models.Tag
	DB.Find(&tags)
	c.JSON(http.StatusOK, tags)
}

func CreateTag(c *gin.Context) {
	var tag models.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := DB.Create(&tag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create tag"})
		return
	}
	c.JSON(http.StatusOK, tag)
}
