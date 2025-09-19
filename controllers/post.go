package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/voduybaokhanh/blog-cms/config"
	"github.com/voduybaokhanh/blog-cms/models"
)

// ======================= CREATE POST =======================
func CreatePost(c *gin.Context) {
	var req struct {
		Title      string `json:"title" binding:"required"`
		Content    string `json:"content" binding:"required"`
		AuthorID   uint   `json:"author_id" binding:"required"`
		CategoryID uint   `json:"category_id"`
		TagIDs     []uint `json:"tag_ids"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{
		Title:      req.Title,
		Content:    req.Content,
		AuthorID:   req.AuthorID,
		CategoryID: req.CategoryID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if len(req.TagIDs) > 0 {
		var tags []models.Tag
		if err := config.DB.Find(&tags, req.TagIDs).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tags"})
			return
		}
		post.Tags = tags
	}

	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, models.MapPostToResponse(post))
}

// ======================= GET ONE POST =======================
func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := config.DB.Preload("Author").
		Preload("Category").
		Preload("Tags").
		First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, models.MapPostToResponse(post))
}

// ======================= GET ALL POSTS =======================
func GetPosts(c *gin.Context) {
	// query params
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	search := c.Query("search")
	category := c.Query("category")
	tags := c.Query("tag") // ví dụ ?tag=1,2

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	var posts []models.Post
	query := config.DB.Preload("Author").Preload("Category").Preload("Tags")

	if search != "" {
		query = query.Where("title LIKE ? OR content LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if category != "" {
		query = query.Where("category_id = ?", category)
	}

	if tags != "" {
		tagIDs := strings.Split(tags, ",")
		query = query.Joins("JOIN post_tags ON post_tags.post_id = posts.id").
			Where("post_tags.tag_id IN ?", tagIDs).
			Group("posts.id")
	}

	// đếm total
	var total int64
	query.Model(&models.Post{}).Count(&total)

	// lấy data + phân trang
	offset := (page - 1) * limit
	if err := query.Limit(limit).Offset(offset).Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	// map sang response DTO
	var resp []models.PostResponse
	for _, p := range posts {
		resp = append(resp, models.MapPostToResponse(p))
	}

	// trả về data + meta pagination
	c.JSON(http.StatusOK, gin.H{
		"data":  resp,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

// ======================= UPDATE POST =======================
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Title      string `json:"title"`
		Content    string `json:"content"`
		CategoryID uint   `json:"category_id"`
		TagIDs     []uint `json:"tag_ids"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var post models.Post
	if err := config.DB.Preload("Tags").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	post.Title = req.Title
	post.Content = req.Content
	post.CategoryID = req.CategoryID
	post.UpdatedAt = time.Now()

	if len(req.TagIDs) > 0 {
		var tags []models.Tag
		if err := config.DB.Find(&tags, req.TagIDs).Error; err == nil {
			config.DB.Model(&post).Association("Tags").Replace(&tags)
		}
	}

	if err := config.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	c.JSON(http.StatusOK, models.MapPostToResponse(post))
}

// ======================= DELETE POST =======================
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Post{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}
