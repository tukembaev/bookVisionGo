package models

import (
	"time"
	"database/sql/driver"
	"errors"
)

// ArticleType - enum для типов статей
type ArticleType string

const (
	ArticleTypeShouldRead  ArticleType = "shouldRead"
	ArticleTypeAnalysis    ArticleType = "analysis"
	ArticleTypeReview      ArticleType = "review"
	ArticleTypeCollection  ArticleType = "collection"
	ArticleTypeGuide       ArticleType = "guide"
	ArticleTypeComparison  ArticleType = "comparison"
	ArticleTypeDiscussion  ArticleType = "discussion"
)

// Value - реализация driver.Valuer для PostgreSQL
func (at ArticleType) Value() (driver.Value, error) {
	return string(at), nil
}

// Scan - реализация sql.Scanner для PostgreSQL
func (at *ArticleType) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return errors.New("cannot scan non-string value into ArticleType")
	}
	*at = ArticleType(str)
	return nil
}

// ArticleReadiness - enum для готовности к чтению
type ArticleReadiness string

const (
	ArticleReadinessMust  ArticleReadiness = "must"
	ArticleReadinessMaybe ArticleReadiness = "maybe"
	ArticleReadinessNo    ArticleReadiness = "no"
)

// Value - реализация driver.Valuer для PostgreSQL
func (ar ArticleReadiness) Value() (driver.Value, error) {
	return string(ar), nil
}

// Scan - реализация sql.Scanner для PostgreSQL
func (ar *ArticleReadiness) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return errors.New("cannot scan non-string value into ArticleReadiness")
	}
	*ar = ArticleReadiness(str)
	return nil
}

// ContentBlockType - enum для типов контент-блоков
type ContentBlockType string

const (
	ContentBlockTypeH2    ContentBlockType = "h2"
	ContentBlockTypeH3    ContentBlockType = "h3"
	ContentBlockTypeP     ContentBlockType = "p"
	ContentBlockTypeQuote ContentBlockType = "quote"
)

// Value - реализация driver.Valuer для PostgreSQL
func (cbt ContentBlockType) Value() (driver.Value, error) {
	return string(cbt), nil
}

// Scan - реализация sql.Scanner для PostgreSQL
func (cbt *ContentBlockType) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return errors.New("cannot scan non-string value into ContentBlockType")
	}
	*cbt = ContentBlockType(str)
	return nil
}

// Article - модель статьи
type Article struct {
	ID                   string           `json:"id" db:"id"`
	Title                string           `json:"title" db:"title"`
	Type                 ArticleType      `json:"type" db:"type"`
	AuthorID             *string          `json:"author_id" db:"author_id"`
	BookID               *string          `json:"book_id" db:"book_id"`
	Excerpt              string           `json:"excerpt" db:"excerpt"`
	CreatedAt            time.Time        `json:"created_at" db:"created_at"`
	Likes                int              `json:"likes" db:"likes"`
	Views                int              `json:"views" db:"views"`
	ReadingMinutes       *int             `json:"reading_minutes" db:"reading_minutes"`
	CoverURL             *string          `json:"cover_url" db:"cover_url"`
	Verified             bool             `json:"verified" db:"verified"`
	VerificationType     *VerificationType `json:"verification_type" db:"verification_type"`
	NoSpoilers           bool             `json:"no_spoilers" db:"no_spoilers"`
	ShouldReadReadiness  ArticleReadiness `json:"should_read_readiness" db:"should_read_readiness"`
}

// ArticleContentBlock - модель контент-блока статьи
type ArticleContentBlock struct {
	ID        string          `json:"id" db:"id"`
	ArticleID string          `json:"article_id" db:"article_id"`
	BlockType ContentBlockType `json:"block_type" db:"block_type"`
	Text      string          `json:"text" db:"text"`
	OrderNum  int             `json:"order_num" db:"order_num"`
	BlockID   *string         `json:"block_id" db:"block_id"`
}

// CreateArticleRequest - DTO для создания статьи
type CreateArticleRequest struct {
	Title               string           `json:"title" binding:"required,max=255"`
	Type                ArticleType      `json:"type" binding:"required"`
	BookID              *string          `json:"book_id"`
	Excerpt             string           `json:"excerpt" binding:"required"`
	ReadingMinutes      *int             `json:"reading_minutes"`
	CoverURL            *string          `json:"cover_url"`
	NoSpoilers          bool             `json:"no_spoilers"`
	ShouldReadReadiness ArticleReadiness `json:"should_read_readiness"`
	ContentBlocks      []CreateContentBlockRequest `json:"content_blocks" binding:"required,min=1"`
}

// CreateContentBlockRequest - DTO для создания контент-блока
type CreateContentBlockRequest struct {
	BlockType ContentBlockType `json:"block_type" binding:"required"`
	Text      string           `json:"text" binding:"required"`
	BlockID   *string         `json:"block_id"`
}

// UpdateArticleRequest - DTO для обновления статьи
type UpdateArticleRequest struct {
	Title               *string          `json:"title"`
	Type                *ArticleType     `json:"type"`
	Excerpt             *string          `json:"excerpt"`
	ReadingMinutes      *int             `json:"reading_minutes"`
	CoverURL            *string          `json:"cover_url"`
	Verified            *bool            `json:"verified"`
	VerificationType    *VerificationType `json:"verification_type"`
	NoSpoilers          *bool            `json:"no_spoilers"`
	ShouldReadReadiness *ArticleReadiness `json:"should_read_readiness"`
}

// ArticleResponse - DTO для ответа API
type ArticleResponse struct {
	ID                   string           `json:"id"`
	Title                string           `json:"title"`
	Type                 ArticleType      `json:"type"`
	AuthorID             *string          `json:"author_id"`
	BookID               *string          `json:"book_id"`
	Excerpt              string           `json:"excerpt"`
	CreatedAt            time.Time        `json:"created_at"`
	Likes                int              `json:"likes"`
	Views                int              `json:"views"`
	ReadingMinutes       *int             `json:"reading_minutes"`
	CoverURL             *string          `json:"cover_url"`
	Verified             bool             `json:"verified"`
	VerificationType     *VerificationType `json:"verification_type"`
	NoSpoilers           bool             `json:"no_spoilers"`
	ShouldReadReadiness  ArticleReadiness `json:"should_read_readiness"`
}

// ArticleContentBlockResponse - DTO для ответа API контент-блока
type ArticleContentBlockResponse struct {
	ID        string          `json:"id"`
	ArticleID string          `json:"article_id"`
	BlockType ContentBlockType `json:"block_type"`
	Text      string          `json:"text"`
	OrderNum  int             `json:"order_num"`
	BlockID   *string         `json:"block_id"`
}

// ToResponse - конвертация Article в ArticleResponse
func (a *Article) ToResponse() *ArticleResponse {
	return &ArticleResponse{
		ID:                   a.ID,
		Title:                a.Title,
		Type:                 a.Type,
		AuthorID:             a.AuthorID,
		BookID:               a.BookID,
		Excerpt:              a.Excerpt,
		CreatedAt:            a.CreatedAt,
		Likes:                a.Likes,
		Views:                a.Views,
		ReadingMinutes:       a.ReadingMinutes,
		CoverURL:             a.CoverURL,
		Verified:             a.Verified,
		VerificationType:     a.VerificationType,
		NoSpoilers:           a.NoSpoilers,
		ShouldReadReadiness:  a.ShouldReadReadiness,
	}
}

// ToResponse - конвертация ArticleContentBlock в ArticleContentBlockResponse
func (acb *ArticleContentBlock) ToResponse() *ArticleContentBlockResponse {
	return &ArticleContentBlockResponse{
		ID:        acb.ID,
		ArticleID: acb.ArticleID,
		BlockType: acb.BlockType,
		Text:      acb.Text,
		OrderNum:  acb.OrderNum,
		BlockID:   acb.BlockID,
	}
}
