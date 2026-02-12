package models

import (
	"time"
)

// Review - модель отзыва
type Review struct {
	ID                 string   `json:"id" db:"id"`
	UserID             string   `json:"user_id" db:"user_id"`
	BookID             string   `json:"book_id" db:"book_id"`
	Rating             int      `json:"rating" db:"rating"`
	Text               string   `json:"text" db:"text"`
	LikedCharacters    []string `json:"liked_characters" db:"liked_characters"`
	DislikedCharacters []string `json:"disliked_characters" db:"disliked_characters"`
	BestParts          []string `json:"best_parts" db:"best_parts"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
}

// Comment - модель комментария
type Comment struct {
	ID             string    `json:"id" db:"id"`
	UserID         string    `json:"user_id" db:"user_id"`
	BookID         *string   `json:"book_id" db:"book_id"`
	PartID         *string   `json:"part_id" db:"part_id"`
	Text           string    `json:"text" db:"text"`
	Likes          int       `json:"likes" db:"likes"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	ParentCommentID *string  `json:"parent_comment_id" db:"parent_comment_id"`
	ReplyToUserID  *string   `json:"reply_to_user_id" db:"reply_to_user_id"`
}

// CreateReviewRequest - DTO для создания отзыва
type CreateReviewRequest struct {
	BookID             string   `json:"book_id" binding:"required"`
	Rating             int      `json:"rating" binding:"required,min=1,max=10"`
	Text               string   `json:"text" binding:"required"`
	LikedCharacters    []string `json:"liked_characters"`
	DislikedCharacters []string `json:"disliked_characters"`
	BestParts          []string `json:"best_parts"`
}

// UpdateReviewRequest - DTO для обновления отзыва
type UpdateReviewRequest struct {
	Rating             *int     `json:"rating"`
	Text               *string  `json:"text"`
	LikedCharacters    []string `json:"liked_characters"`
	DislikedCharacters []string `json:"disliked_characters"`
	BestParts          []string `json:"best_parts"`
}

// CreateCommentRequest - DTO для создания комментария
type CreateCommentRequest struct {
	BookID         *string `json:"book_id"`
	PartID         *string `json:"part_id"`
	Text           string  `json:"text" binding:"required"`
	ParentCommentID *string `json:"parent_comment_id"`
	ReplyToUserID  *string `json:"reply_to_user_id"`
}

// UpdateCommentRequest - DTO для обновления комментария
type UpdateCommentRequest struct {
	Text *string `json:"text" binding:"required"`
}

// ReviewResponse - DTO для ответа API
type ReviewResponse struct {
	ID                 string   `json:"id"`
	UserID             string   `json:"user_id"`
	BookID             string   `json:"book_id"`
	Rating             int      `json:"rating"`
	Text               string   `json:"text"`
	LikedCharacters    []string `json:"liked_characters"`
	DislikedCharacters []string `json:"disliked_characters"`
	BestParts          []string `json:"best_parts"`
	CreatedAt          time.Time `json:"created_at"`
}

// CommentResponse - DTO для ответа API
type CommentResponse struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	BookID         *string   `json:"book_id"`
	PartID         *string   `json:"part_id"`
	Text           string    `json:"text"`
	Likes          int       `json:"likes"`
	CreatedAt      time.Time `json:"created_at"`
	ParentCommentID *string  `json:"parent_comment_id"`
	ReplyToUserID  *string   `json:"reply_to_user_id"`
}

// ToResponse - конвертация Review в ReviewResponse
func (r *Review) ToResponse() *ReviewResponse {
	return &ReviewResponse{
		ID:                 r.ID,
		UserID:             r.UserID,
		BookID:             r.BookID,
		Rating:             r.Rating,
		Text:               r.Text,
		LikedCharacters:    r.LikedCharacters,
		DislikedCharacters: r.DislikedCharacters,
		BestParts:          r.BestParts,
		CreatedAt:          r.CreatedAt,
	}
}

// ToResponse - конвертация Comment в CommentResponse
func (c *Comment) ToResponse() *CommentResponse {
	return &CommentResponse{
		ID:             c.ID,
		UserID:         c.UserID,
		BookID:         c.BookID,
		PartID:         c.PartID,
		Text:           c.Text,
		Likes:          c.Likes,
		CreatedAt:      c.CreatedAt,
		ParentCommentID: c.ParentCommentID,
		ReplyToUserID:  c.ReplyToUserID,
	}
}
