package models

import "time"

type Post struct {
	ID         uint `gorm:"primaryKey"`
	Title      string
	Content    string
	AuthorID   uint
	Author     User `gorm:"foreignKey:AuthorID"`
	CategoryID uint
	Category   Category `gorm:"foreignKey:CategoryID"`
	Tags       []Tag    `gorm:"many2many:post_tags;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// DTO trả về
type PostResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Category  string    `json:"category"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MapPostToResponse(p Post) PostResponse {
	tags := make([]string, len(p.Tags))
	for i, t := range p.Tags {
		tags[i] = t.Name
	}

	return PostResponse{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		Author:    p.Author.Name,
		Category:  p.Category.Name,
		Tags:      tags,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
