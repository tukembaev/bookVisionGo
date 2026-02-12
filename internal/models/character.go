package models

import (
	"database/sql/driver"
	"errors"
)

// CharacterSource - enum для источника персонажа
type CharacterSource string

const (
	CharacterSourceWiki      CharacterSource = "wiki"
	CharacterSourceCommunity CharacterSource = "community"
)

// Value - реализация driver.Valuer для PostgreSQL
func (cs CharacterSource) Value() (driver.Value, error) {
	return string(cs), nil
}

// Scan - реализация sql.Scanner для PostgreSQL
func (cs *CharacterSource) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return errors.New("cannot scan non-string value into CharacterSource")
	}
	*cs = CharacterSource(str)
	return nil
}

// Character - модель персонажа
type Character struct {
	ID              string          `json:"id" db:"id"`
	BookID          string          `json:"book_id" db:"book_id"`
	Name            string          `json:"name" db:"name"`
	Description     string          `json:"description" db:"description"`
	Source          CharacterSource `json:"source" db:"source"`
	Verified        bool            `json:"verified" db:"verified"`
	PopularityScore *int            `json:"popularity_score" db:"popularity_score"`
}

// CharacterProfile - расширенный профиль персонажа
type CharacterProfile struct {
	ID                    string   `json:"id" db:"id"`
	Aliases               []string `json:"aliases" db:"aliases"`
	ImageURL              *string  `json:"image_url" db:"image_url"`
	Age                   *string  `json:"age" db:"age"`
	Height                *string  `json:"height" db:"height"`
	Weight                *string  `json:"weight" db:"weight"`
	SocialStatus          *string  `json:"social_status" db:"social_status"`
	DescriptionNoSpoilers string   `json:"description_no_spoilers" db:"description_no_spoilers"`
	DescriptionSpoilers   string   `json:"description_spoilers" db:"description_spoilers"`
	QuotesNoSpoilers      []string `json:"quotes_no_spoilers" db:"quotes_no_spoilers"`
	QuotesSpoilers        []string `json:"quotes_spoilers" db:"quotes_spoilers"`
}

// CharacterIllustration - иллюстрация персонажа
type CharacterIllustration struct {
	ID          string `json:"id" db:"id"`
	CharacterID string `json:"character_id" db:"character_id"`
	ImageURL    string `json:"image_url" db:"image_url"`
	AuthorName  string `json:"author_name" db:"author_name"`
}

// CreateCharacterRequest - DTO для создания персонажа
type CreateCharacterRequest struct {
	BookID          string          `json:"book_id" binding:"required"`
	Name            string          `json:"name" binding:"required,max=255"`
	Description     string          `json:"description" binding:"required"`
	Source          CharacterSource `json:"source"`
	PopularityScore *int            `json:"popularity_score"`
}

// UpdateCharacterRequest - DTO для обновления персонажа
type UpdateCharacterRequest struct {
	Name            *string          `json:"name"`
	Description     *string          `json:"description"`
	Source          *CharacterSource `json:"source"`
	Verified        *bool            `json:"verified"`
	PopularityScore *int             `json:"popularity_score"`
}

// CreateCharacterProfileRequest - DTO для создания профиля персонажа
type CreateCharacterProfileRequest struct {
	Aliases               []string `json:"aliases"`
	ImageURL              *string  `json:"image_url"`
	Age                   *string  `json:"age"`
	Height                *string  `json:"height"`
	Weight                *string  `json:"weight"`
	SocialStatus          *string  `json:"social_status"`
	DescriptionNoSpoilers string   `json:"description_no_spoilers" binding:"required"`
	DescriptionSpoilers   string   `json:"description_spoilers" binding:"required"`
	QuotesNoSpoilers      []string `json:"quotes_no_spoilers"`
	QuotesSpoilers        []string `json:"quotes_spoilers"`
}

// CharacterResponse - DTO для ответа API
type CharacterResponse struct {
	ID              string          `json:"id"`
	BookID          string          `json:"book_id"`
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	Source          CharacterSource `json:"source"`
	Verified        bool            `json:"verified"`
	PopularityScore *int            `json:"popularity_score"`
}

// CharacterProfileResponse - DTO для ответа API профиля
type CharacterProfileResponse struct {
	ID                    string   `json:"id"`
	Aliases               []string `json:"aliases"`
	ImageURL              *string  `json:"image_url"`
	Age                   *string  `json:"age"`
	Height                *string  `json:"height"`
	Weight                *string  `json:"weight"`
	SocialStatus          *string  `json:"social_status"`
	DescriptionNoSpoilers string   `json:"description_no_spoilers"`
	DescriptionSpoilers   string   `json:"description_spoilers"`
	QuotesNoSpoilers      []string `json:"quotes_no_spoilers"`
	QuotesSpoilers        []string `json:"quotes_spoilers"`
}

// CharacterIllustrationResponse - DTO для ответа API иллюстрации
type CharacterIllustrationResponse struct {
	ID          string `json:"id"`
	CharacterID string `json:"character_id"`
	ImageURL    string `json:"image_url"`
	AuthorName  string `json:"author_name"`
}

// ToResponse - конвертация Character в CharacterResponse
func (c *Character) ToResponse() *CharacterResponse {
	return &CharacterResponse{
		ID:              c.ID,
		BookID:          c.BookID,
		Name:            c.Name,
		Description:     c.Description,
		Source:          c.Source,
		Verified:        c.Verified,
		PopularityScore: c.PopularityScore,
	}
}

// ToResponse - конвертация CharacterProfile в CharacterProfileResponse
func (cp *CharacterProfile) ToResponse() *CharacterProfileResponse {
	return &CharacterProfileResponse{
		ID:                    cp.ID,
		Aliases:               cp.Aliases,
		ImageURL:              cp.ImageURL,
		Age:                   cp.Age,
		Height:                cp.Height,
		Weight:                cp.Weight,
		SocialStatus:          cp.SocialStatus,
		DescriptionNoSpoilers: cp.DescriptionNoSpoilers,
		DescriptionSpoilers:   cp.DescriptionSpoilers,
		QuotesNoSpoilers:      cp.QuotesNoSpoilers,
		QuotesSpoilers:        cp.QuotesSpoilers,
	}
}

// ToResponse - конвертация CharacterIllustration в CharacterIllustrationResponse
func (ci *CharacterIllustration) ToResponse() *CharacterIllustrationResponse {
	return &CharacterIllustrationResponse{
		ID:          ci.ID,
		CharacterID: ci.CharacterID,
		ImageURL:    ci.ImageURL,
		AuthorName:  ci.AuthorName,
	}
}
