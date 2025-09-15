package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/voduybaokhanh/blog-cms/models"
)

func GetCategories(c *gin.Context) {
	var categories []models.Category
	DB.Find(&categories)
	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create category"})
		return
	}
	c.JSON(http.StatusOK, category)
}
