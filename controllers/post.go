package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/voduybaokhanh/blog-cms/models"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := DB.Preload("Author").Preload("Category").Preload("Tags").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot fetch posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	var post models.Post
	if err := DB.Preload("Author").Preload("Category").Preload("Tags").First(&post, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func CreatePost(c *gin.Context) {
	var input struct {
		Title      string `json:"title" binding:"required"`
		Content    string `json:"content" binding:"required"`
		CategoryID uint   `json:"category_id"`
		TagIDs     []uint `json:"tag_ids"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")

	post := models.Post{
		Title:      input.Title,
		Content:    input.Content,
		AuthorID:   userID.(uint),
		CategoryID: input.CategoryID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if len(input.TagIDs) > 0 {
		var tags []models.Tag
		DB.Find(&tags, input.TagIDs)
		post.Tags = tags
	}

	if err := DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create post"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	if err := DB.First(&post, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	var input struct {
		Title      string `json:"title"`
		Content    string `json:"content"`
		CategoryID uint   `json:"category_id"`
		TagIDs     []uint `json:"tag_ids"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.Title = input.Title
	post.Content = input.Content
	post.CategoryID = input.CategoryID
	post.UpdatedAt = time.Now()

	if len(input.TagIDs) > 0 {
		var tags []models.Tag
		DB.Find(&tags, input.TagIDs)
		DB.Model(&post).Association("Tags").Replace(tags)
	}

	if err := DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot update post"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func DeletePost(c *gin.Context) {
	if err := DB.Delete(&models.Post{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot delete post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}
