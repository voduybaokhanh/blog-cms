package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/voduybaokhanh/blog-cms/models"
	"gorm.io/gorm"
)

func SetPostDB(db *gorm.DB) {
	DB = db
}

// ------------------ CREATE ------------------
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

	// lấy tags từ DB
	var tags []models.Tag
	if len(input.TagIDs) > 0 {
		DB.Where("id IN ?", input.TagIDs).Find(&tags)
	}

	var authorID uint
	switch v := userID.(type) {
	case uint:
		authorID = v
	case float64:
		authorID = uint(v)
	default:
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	post := models.Post{
		Title:      input.Title,
		Content:    input.Content,
		AuthorID:   authorID,
		CategoryID: input.CategoryID,
		Tags:       tags,
	}

	if err := DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	DB.Preload("Author").Preload("Category").Preload("Tags").First(&post, post.ID)
	c.JSON(http.StatusOK, models.MapPostToResponse(post))
}

// ------------------ READ (ALL) ------------------
func GetPosts(c *gin.Context) {
	var posts []models.Post
	DB.Preload("Author").Preload("Category").Preload("Tags").Find(&posts)

	responses := make([]models.PostResponse, len(posts))
	for i, p := range posts {
		responses[i] = models.MapPostToResponse(p)
	}

	c.JSON(http.StatusOK, responses)
}

// ------------------ READ (ONE) ------------------
func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := DB.Preload("Author").Preload("Category").Preload("Tags").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, models.MapPostToResponse(post))
}

// ------------------ UPDATE ------------------
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := DB.Preload("Tags").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")

	// chỉ admin hoặc chính author mới được sửa
	if role != "admin" {
		var uid uint
		switch v := userID.(type) {
		case uint:
			uid = v
		case float64:
			uid = uint(v)
		default:
			c.JSON(http.StatusForbidden, gin.H{"error": "Not allowed"})
			return
		}
		if uid != post.AuthorID {
			c.JSON(http.StatusForbidden, gin.H{"error": "Not allowed"})
			return
		}
	}

	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		TagIDs  []uint `json:"tag_ids"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Title != "" {
		post.Title = input.Title
	}
	if input.Content != "" {
		post.Content = input.Content
	}

	// update tags nếu có
	if input.TagIDs != nil {
		var tags []models.Tag
		DB.Where("id IN ?", input.TagIDs).Find(&tags)
		post.Tags = tags
	}

	DB.Save(&post)
	DB.Preload("Author").Preload("Category").Preload("Tags").First(&post, post.ID)

	c.JSON(http.StatusOK, models.MapPostToResponse(post))
}

// ------------------ DELETE ------------------
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")

	// chỉ admin hoặc chính author mới được xoá
	if role != "admin" {
		var uid uint
		switch v := userID.(type) {
		case uint:
			uid = v
		case float64:
			uid = uint(v)
		default:
			c.JSON(http.StatusForbidden, gin.H{"error": "Not allowed"})
			return
		}
		if uid != post.AuthorID {
			c.JSON(http.StatusForbidden, gin.H{"error": "Not allowed"})
			return
		}
	}

	DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}
