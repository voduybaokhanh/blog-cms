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

func DeleteTag(c *gin.Context) {
	id := c.Param("id")
	var tag models.Tag
	if err := DB.First(&tag, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}
	if err := DB.Delete(&tag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot delete tag"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tag deleted"})
}
