package models

import (
	"database/sql/driver"
	"errors"
)

// PlaylistCreator - enum для создателя плейлиста
type PlaylistCreator string

const (
	PlaylistCreatorSystem PlaylistCreator = "system"
	PlaylistCreatorUser   PlaylistCreator = "user"
)

// Value - реализация driver.Valuer для PostgreSQL
func (pc PlaylistCreator) Value() (driver.Value, error) {
	return string(pc), nil
}

// Scan - реализация sql.Scanner для PostgreSQL
func (pc *PlaylistCreator) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return errors.New("cannot scan non-string value into PlaylistCreator")
	}
	*pc = PlaylistCreator(str)
	return nil
}

// PlaylistPlatform - enum для платформ плейлистов
type PlaylistPlatform string

const (
	PlaylistPlatformSpotify PlaylistPlatform = "spotify"
	PlaylistPlatformYouTube PlaylistPlatform = "youtube"
	PlaylistPlatformText    PlaylistPlatform = "text"
)

// Value - реализация driver.Valuer для PostgreSQL
func (pp PlaylistPlatform) Value() (driver.Value, error) {
	return string(pp), nil
}

// Scan - реализация sql.Scanner для PostgreSQL
func (pp *PlaylistPlatform) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return errors.New("cannot scan non-string value into PlaylistPlatform")
	}
	*pp = PlaylistPlatform(str)
	return nil
}

// Playlist - модель плейлиста
type Playlist struct {
	ID        string          `json:"id" db:"id"`
	Title     string          `json:"title" db:"title"`
	MoodTag   string          `json:"mood_tag" db:"mood_tag"`
	Tracks    []string        `json:"tracks" db:"tracks"`
	CreatedBy PlaylistCreator `json:"created_by" db:"created_by"`
	UserID    *string         `json:"user_id" db:"user_id"`
}

// BookPlaylist - модель связи книги с плейлистом
type BookPlaylist struct {
	ID         string           `json:"id" db:"id"`
	BookID     *string          `json:"book_id" db:"book_id"`
	PartID     *string          `json:"part_id" db:"part_id"`
	PlaylistID string           `json:"playlist_id" db:"playlist_id"`
	Platform   PlaylistPlatform `json:"platform" db:"platform"`
	URL        *string          `json:"url" db:"url"`
	CreatedBy  PlaylistCreator  `json:"created_by" db:"created_by"`
}

// UserBookProgress - модель прогресса чтения книги пользователем
type UserBookProgress struct {
	ID               string   `json:"id" db:"id"`
	UserID           string   `json:"user_id" db:"user_id"`
	BookID           string   `json:"book_id" db:"book_id"`
	CompletedPartIDs []string `json:"completed_part_ids" db:"completed_part_ids"`
	CurrentPartID    *string  `json:"current_part_id" db:"current_part_id"`
	IsCompleted      bool     `json:"is_completed" db:"is_completed"`
	CompletedAt      *string  `json:"completed_at" db:"completed_at"`
}

// ReadBookForm - модель формы прочитанной книги
type ReadBookForm struct {
	ID                  string  `json:"id" db:"id"`
	UserID              string  `json:"user_id" db:"user_id"`
	BookID              string  `json:"book_id" db:"book_id"`
	FavoriteCharacterID *string `json:"favorite_character_id" db:"favorite_character_id"`
	DislikedCharacterID *string `json:"disliked_character_id" db:"disliked_character_id"`
	BestPartID          *string `json:"best_part_id" db:"best_part_id"`
	Rating              int     `json:"rating" db:"rating"`
	Thoughts            *string `json:"thoughts" db:"thoughts"`
	CompletedAt         string  `json:"completed_at" db:"completed_at"`
}

// Quote - модель цитаты
type Quote struct {
	ID        string  `json:"id" db:"id"`
	UserID    string  `json:"user_id" db:"user_id"`
	BookID    string  `json:"book_id" db:"book_id"`
	PartID    *string `json:"part_id" db:"part_id"`
	Text      string  `json:"text" db:"text"`
	CreatedAt string  `json:"created_at" db:"created_at"`
}

// CreatePlaylistRequest - DTO для создания плейлиста
type CreatePlaylistRequest struct {
	Title   string   `json:"title" binding:"required,max=255"`
	MoodTag string   `json:"mood_tag" binding:"required,max=50"`
	Tracks  []string `json:"tracks" binding:"required,min=1"`
}

// UpdatePlaylistRequest - DTO для обновления плейлиста
type UpdatePlaylistRequest struct {
	Title   *string  `json:"title"`
	MoodTag *string  `json:"mood_tag"`
	Tracks  []string `json:"tracks"`
}

// CreateBookPlaylistRequest - DTO для связи книги с плейлистом
type CreateBookPlaylistRequest struct {
	BookID     *string          `json:"book_id"`
	PartID     *string          `json:"part_id"`
	PlaylistID string           `json:"playlist_id" binding:"required"`
	Platform   PlaylistPlatform `json:"platform" binding:"required"`
	URL        *string          `json:"url"`
}

// UpdateBookProgressRequest - DTO для обновления прогресса чтения
type UpdateBookProgressRequest struct {
	CompletedPartIDs []string `json:"completed_part_ids"`
	CurrentPartID    *string  `json:"current_part_id"`
	IsCompleted      bool     `json:"is_completed"`
}

// CreateQuoteRequest - DTO для создания цитаты
type CreateQuoteRequest struct {
	BookID string  `json:"book_id" binding:"required"`
	PartID *string `json:"part_id"`
	Text   string  `json:"text" binding:"required,max=1000"`
}

// PlaylistResponse - DTO для ответа API
type PlaylistResponse struct {
	ID        string          `json:"id"`
	Title     string          `json:"title"`
	MoodTag   string          `json:"mood_tag"`
	Tracks    []string        `json:"tracks"`
	CreatedBy PlaylistCreator `json:"created_by"`
	UserID    *string         `json:"user_id"`
}

// BookPlaylistResponse - DTO для ответа API связи книги с плейлистом
type BookPlaylistResponse struct {
	ID         string           `json:"id"`
	BookID     *string          `json:"book_id"`
	PartID     *string          `json:"part_id"`
	PlaylistID string           `json:"playlist_id"`
	Platform   PlaylistPlatform `json:"platform"`
	URL        *string          `json:"url"`
	CreatedBy  PlaylistCreator  `json:"created_by"`
}

// UserBookProgressResponse - DTO для ответа API прогресса
type UserBookProgressResponse struct {
	ID               string   `json:"id"`
	UserID           string   `json:"user_id"`
	BookID           string   `json:"book_id"`
	CompletedPartIDs []string `json:"completed_part_ids"`
	CurrentPartID    *string  `json:"current_part_id"`
	IsCompleted      bool     `json:"is_completed"`
	CompletedAt      *string  `json:"completed_at"`
}

// QuoteResponse - DTO для ответа API цитаты
type QuoteResponse struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	BookID    string  `json:"book_id"`
	PartID    *string `json:"part_id"`
	Text      string  `json:"text"`
	CreatedAt string  `json:"created_at"`
}

// ToResponse - конвертация Playlist в PlaylistResponse
func (p *Playlist) ToResponse() *PlaylistResponse {
	return &PlaylistResponse{
		ID:        p.ID,
		Title:     p.Title,
		MoodTag:   p.MoodTag,
		Tracks:    p.Tracks,
		CreatedBy: p.CreatedBy,
		UserID:    p.UserID,
	}
}

// ToResponse - конвертация BookPlaylist в BookPlaylistResponse
func (bp *BookPlaylist) ToResponse() *BookPlaylistResponse {
	return &BookPlaylistResponse{
		ID:         bp.ID,
		BookID:     bp.BookID,
		PartID:     bp.PartID,
		PlaylistID: bp.PlaylistID,
		Platform:   bp.Platform,
		URL:        bp.URL,
		CreatedBy:  bp.CreatedBy,
	}
}

// ToResponse - конвертация UserBookProgress в UserBookProgressResponse
func (ubp *UserBookProgress) ToResponse() *UserBookProgressResponse {
	return &UserBookProgressResponse{
		ID:               ubp.ID,
		UserID:           ubp.UserID,
		BookID:           ubp.BookID,
		CompletedPartIDs: ubp.CompletedPartIDs,
		CurrentPartID:    ubp.CurrentPartID,
		IsCompleted:      ubp.IsCompleted,
		CompletedAt:      ubp.CompletedAt,
	}
}

// ToResponse - конвертация Quote в QuoteResponse
func (q *Quote) ToResponse() *QuoteResponse {
	return &QuoteResponse{
		ID:        q.ID,
		UserID:    q.UserID,
		BookID:    q.BookID,
		PartID:    q.PartID,
		Text:      q.Text,
		CreatedAt: q.CreatedAt,
	}
}
