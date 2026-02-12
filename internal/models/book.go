package models

import (
	"time"
	"database/sql/driver"
	"errors"
)

// AgeRating - enum для возрастных рейтингов
type AgeRating string

const (
	AgeRating6  AgeRating = "6+"
	AgeRating12 AgeRating = "12+"
	AgeRating16 AgeRating = "16+"
	AgeRating18 AgeRating = "18+"
)

// Value - реализация driver.Valuer для PostgreSQL
func (ar AgeRating) Value() (driver.Value, error) {
	return string(ar), nil
}

// Scan - реализация sql.Scanner для PostgreSQL
func (ar *AgeRating) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return errors.New("cannot scan non-string value into AgeRating")
	}
	*ar = AgeRating(str)
	return nil
}

// VerificationType - enum для типов верификации
type VerificationType string

const (
	VerificationTypeAI        VerificationType = "AI"
	VerificationTypeCommunity VerificationType = "Community"
)

// Value - реализация driver.Valuer для PostgreSQL
func (vt VerificationType) Value() (driver.Value, error) {
	return string(vt), nil
}

// Scan - реализация sql.Scanner для PostgreSQL
func (vt *VerificationType) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return errors.New("cannot scan non-string value into VerificationType")
	}
	*vt = VerificationType(str)
	return nil
}

// Book - модель книги
type Book struct {
	ID             string           `json:"id" db:"id"`
	Title          string           `json:"title" db:"title"`
	OriginalTitle  *string          `json:"original_title" db:"original_title"`
	Author         string           `json:"author" db:"author"`
	Year           *int             `json:"year" db:"year"`
	Genres         []string         `json:"genres" db:"genres"`
	AgeRating      *AgeRating       `json:"age_rating" db:"age_rating"`
	AuthorCountry  *string          `json:"author_country" db:"author_country"`
	Description    string           `json:"description" db:"description"`
	CoverURL       *string          `json:"cover_url" db:"cover_url"`
	PagesCount     int              `json:"pages_count" db:"pages_count"`
	Verified       bool             `json:"verified" db:"verified"`
	VerificationType *VerificationType `json:"verification_type" db:"verification_type"`
	CreatedBy      *string          `json:"created_by" db:"created_by"`
	CreatedAt      time.Time        `json:"created_at" db:"created_at"`
	Tags           []string         `json:"tags" db:"tags"`
	AverageRating  float64          `json:"average_rating" db:"average_rating"`
	RatingCount    int              `json:"rating_count" db:"rating_count"`
}

// BookPart - модель части/главы книги
type BookPart struct {
	ID           string    `json:"id" db:"id"`
	BookID       string    `json:"book_id" db:"book_id"`
	Title        string    `json:"title" db:"title"`
	OrderNum     int       `json:"order_num" db:"order_num"`
	PageStart    *int      `json:"page_start" db:"page_start"`
	PageEnd      *int      `json:"page_end" db:"page_end"`
	MoodTags     []string  `json:"mood_tags" db:"mood_tags"`
	AverageRating *float64 `json:"average_rating" db:"average_rating"`
}

// CreateBookRequest - DTO для создания книги
type CreateBookRequest struct {
	Title          string     `json:"title" binding:"required,max=255"`
	OriginalTitle  *string    `json:"original_title"`
	Author         string     `json:"author" binding:"required,max=255"`
	Year           *int       `json:"year"`
	Genres         []string   `json:"genres"`
	AgeRating      *AgeRating `json:"age_rating"`
	AuthorCountry  *string    `json:"author_country"`
	Description    string     `json:"description" binding:"required"`
	CoverURL       *string    `json:"cover_url"`
	PagesCount     int        `json:"pages_count" binding:"required,min=1"`
	Tags           []string   `json:"tags"`
}

// UpdateBookRequest - DTO для обновления книги
type UpdateBookRequest struct {
	Title          *string    `json:"title"`
	OriginalTitle  *string    `json:"original_title"`
	Author         *string    `json:"author"`
	Year           *int       `json:"year"`
	Genres         []string   `json:"genres"`
	AgeRating      *AgeRating `json:"age_rating"`
	AuthorCountry  *string    `json:"author_country"`
	Description    *string    `json:"description"`
	CoverURL       *string    `json:"cover_url"`
	PagesCount     *int       `json:"pages_count"`
	Tags           []string   `json:"tags"`
	Verified       *bool      `json:"verified"`
	VerificationType *VerificationType `json:"verification_type"`
}

// BookResponse - DTO для ответа API
type BookResponse struct {
	ID             string           `json:"id"`
	Title          string           `json:"title"`
	OriginalTitle  *string          `json:"original_title"`
	Author         string           `json:"author"`
	Year           *int             `json:"year"`
	Genres         []string         `json:"genres"`
	AgeRating      *AgeRating       `json:"age_rating"`
	AuthorCountry  *string          `json:"author_country"`
	Description    string           `json:"description"`
	CoverURL       *string          `json:"cover_url"`
	PagesCount     int              `json:"pages_count"`
	Verified       bool             `json:"verified"`
	VerificationType *VerificationType `json:"verification_type"`
	CreatedBy      *string          `json:"created_by"`
	CreatedAt      time.Time        `json:"created_at"`
	Tags           []string         `json:"tags"`
	AverageRating  float64          `json:"average_rating"`
	RatingCount    int              `json:"rating_count"`
}

// BookPartResponse - DTO для ответа API части книги
type BookPartResponse struct {
	ID           string    `json:"id"`
	BookID       string    `json:"book_id"`
	Title        string    `json:"title"`
	OrderNum     int       `json:"order_num"`
	PageStart    *int      `json:"page_start"`
	PageEnd      *int      `json:"page_end"`
	MoodTags     []string  `json:"mood_tags"`
	AverageRating *float64 `json:"average_rating"`
}

// ToResponse - конвертация Book в BookResponse
func (b *Book) ToResponse() *BookResponse {
	return &BookResponse{
		ID:             b.ID,
		Title:          b.Title,
		OriginalTitle:  b.OriginalTitle,
		Author:         b.Author,
		Year:           b.Year,
		Genres:         b.Genres,
		AgeRating:      b.AgeRating,
		AuthorCountry:  b.AuthorCountry,
		Description:    b.Description,
		CoverURL:       b.CoverURL,
		PagesCount:     b.PagesCount,
		Verified:       b.Verified,
		VerificationType: b.VerificationType,
		CreatedBy:      b.CreatedBy,
		CreatedAt:      b.CreatedAt,
		Tags:           b.Tags,
		AverageRating:  b.AverageRating,
		RatingCount:    b.RatingCount,
	}
}

// ToResponse - конвертация BookPart в BookPartResponse
func (bp *BookPart) ToResponse() *BookPartResponse {
	return &BookPartResponse{
		ID:            bp.ID,
		BookID:        bp.BookID,
		Title:         bp.Title,
		OrderNum:      bp.OrderNum,
		PageStart:     bp.PageStart,
		PageEnd:       bp.PageEnd,
		MoodTags:      bp.MoodTags,
		AverageRating: bp.AverageRating,
	}
}
