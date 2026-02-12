package models

import (
	"time"
	"database/sql/driver"
	"errors"
)

// ChallengeType - enum для типов челленджей
type ChallengeType string

const (
	ChallengeTypeBooks   ChallengeType = "books"
	ChallengeTypeReviews ChallengeType = "reviews"
)

// Value - реализация driver.Valuer для PostgreSQL
func (ct ChallengeType) Value() (driver.Value, error) {
	return string(ct), nil
}

// Scan - реализация sql.Scanner для PostgreSQL
func (ct *ChallengeType) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return errors.New("cannot scan non-string value into ChallengeType")
	}
	*ct = ChallengeType(str)
	return nil
}

// ChallengeStatus - enum для статусов челленджей
type ChallengeStatus string

const (
	ChallengeStatusActive    ChallengeStatus = "active"
	ChallengeStatusCompleted ChallengeStatus = "completed"
)

// Value - реализация driver.Valuer для PostgreSQL
func (cs ChallengeStatus) Value() (driver.Value, error) {
	return string(cs), nil
}

// Scan - реализация sql.Scanner для PostgreSQL
func (cs *ChallengeStatus) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return errors.New("cannot scan non-string value into ChallengeStatus")
	}
	*cs = ChallengeStatus(str)
	return nil
}

// Challenge - модель челленджа
type Challenge struct {
	ID           string        `json:"id" db:"id"`
	Title        string        `json:"title" db:"title"`
	Description  string        `json:"description" db:"description"`
	Type         ChallengeType `json:"type" db:"type"`
	TargetCount  int           `json:"target_count" db:"target_count"`
	RewardPoints int           `json:"reward_points" db:"reward_points"`
	CreatedAt    time.Time     `json:"created_at" db:"created_at"`
}

// UserChallengeProgress - модель прогресса пользователя в челлендже
type UserChallengeProgress struct {
	ID           string          `json:"id" db:"id"`
	UserID       string          `json:"user_id" db:"user_id"`
	ChallengeID  string          `json:"challenge_id" db:"challenge_id"`
	Status       ChallengeStatus `json:"status" db:"status"`
	ProgressCount int            `json:"progress_count" db:"progress_count"`
	StartedAt    time.Time       `json:"started_at" db:"started_at"`
	CompletedAt  *time.Time      `json:"completed_at" db:"completed_at"`
}

// CreateChallengeRequest - DTO для создания челленджа
type CreateChallengeRequest struct {
	Title        string        `json:"title" binding:"required,max=255"`
	Description  string        `json:"description" binding:"required"`
	Type         ChallengeType `json:"type" binding:"required"`
	TargetCount  int           `json:"target_count" binding:"required,min=1"`
	RewardPoints int           `json:"reward_points" binding:"required,min=1"`
}

// UpdateChallengeRequest - DTO для обновления челленджа
type UpdateChallengeRequest struct {
	Title        *string        `json:"title"`
	Description  *string        `json:"description"`
	Type         *ChallengeType `json:"type"`
	TargetCount  *int           `json:"target_count"`
	RewardPoints *int           `json:"reward_points"`
}

// JoinChallengeRequest - DTO для присоединения к челленджу
type JoinChallengeRequest struct {
	ChallengeID string `json:"challenge_id" binding:"required"`
}

// UpdateProgressRequest - DTO для обновления прогресса
type UpdateProgressRequest struct {
	ProgressCount int `json:"progress_count" binding:"required,min=0"`
}

// ChallengeResponse - DTO для ответа API
type ChallengeResponse struct {
	ID           string        `json:"id"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Type         ChallengeType `json:"type"`
	TargetCount  int           `json:"target_count"`
	RewardPoints int           `json:"reward_points"`
	CreatedAt    time.Time     `json:"created_at"`
}

// UserChallengeProgressResponse - DTO для ответа API прогресса
type UserChallengeProgressResponse struct {
	ID            string          `json:"id"`
	UserID        string          `json:"user_id"`
	ChallengeID   string          `json:"challenge_id"`
	Status        ChallengeStatus `json:"status"`
	ProgressCount int             `json:"progress_count"`
	StartedAt     time.Time       `json:"started_at"`
	CompletedAt   *time.Time      `json:"completed_at"`
}

// ToResponse - конвертация Challenge в ChallengeResponse
func (c *Challenge) ToResponse() *ChallengeResponse {
	return &ChallengeResponse{
		ID:           c.ID,
		Title:        c.Title,
		Description:  c.Description,
		Type:         c.Type,
		TargetCount:  c.TargetCount,
		RewardPoints: c.RewardPoints,
		CreatedAt:    c.CreatedAt,
	}
}

// ToResponse - конвертация UserChallengeProgress в UserChallengeProgressResponse
func (ucp *UserChallengeProgress) ToResponse() *UserChallengeProgressResponse {
	return &UserChallengeProgressResponse{
		ID:            ucp.ID,
		UserID:        ucp.UserID,
		ChallengeID:   ucp.ChallengeID,
		Status:        ucp.Status,
		ProgressCount: ucp.ProgressCount,
		StartedAt:     ucp.StartedAt,
		CompletedAt:   ucp.CompletedAt,
	}
}
