package models

import (
	"database/sql/driver"
	"errors"
	"time"
)

// UserRole - enum для ролей пользователя
type UserRole string

const (
	UserRoleUser      UserRole = "user"
	UserRoleModerator UserRole = "moderator"
	UserRoleAdmin     UserRole = "admin"
)

// Value - реализация driver.Valuer для PostgreSQL
func (ur UserRole) Value() (driver.Value, error) {
	return string(ur), nil
}

// Scan - реализация sql.Scanner для PostgreSQL
func (ur *UserRole) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return errors.New("cannot scan non-string value into UserRole")
	}
	*ur = UserRole(str)
	return nil
}

// User - модель пользователя
type User struct {
	ID                 string    `json:"id" db:"id"`
	Username           string    `json:"username" db:"username"`
	Email              string    `json:"email" db:"email"`
	PasswordHash       string    `json:"-" db:"password_hash"` // "-" не включать в JSON
	AvatarURL          *string   `json:"avatar_url" db:"avatar_url"`
	Role               UserRole  `json:"role" db:"role"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	BooksRead          int       `json:"books_read" db:"books_read"`
	ReviewsCount       int       `json:"reviews_count" db:"reviews_count"`
	LikesReceived      int       `json:"likes_received" db:"likes_received"`
	ProfileVisibility  string    `json:"profile_visibility" db:"profile_visibility"`
	ActivityVisibility string    `json:"activity_visibility" db:"activity_visibility"`
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50" example:"arif123"`
	Email    string `json:"email" binding:"required,email" example:"arif@example.com"`
	Password string `json:"password" binding:"required,min=6" example:"StrongP@ssw0rd"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"arif"`
	Password string `json:"password" binding:"required" example:"arif123"`
}

type UpdateUserRequest struct {
	Username  *string   `json:"username" example:"arif123"`
	AvatarURL *string   `json:"avatar_url" example:"https://example.com/avatar.png"`
	Role      *UserRole `json:"role" example:"user"`
}

// UserResponse - DTO для ответа API (без пароля)
type UserResponse struct {
	ID                 string    `json:"id"`
	Username           string    `json:"username"`
	Email              string    `json:"email"`
	AvatarURL          *string   `json:"avatar_url"`
	Role               UserRole  `json:"role"`
	CreatedAt          time.Time `json:"created_at"`
	BooksRead          int       `json:"books_read"`
	ReviewsCount       int       `json:"reviews_count"`
	LikesReceived      int       `json:"likes_received"`
	ProfileVisibility  string    `json:"profile_visibility"`
	ActivityVisibility string    `json:"activity_visibility"`
}

// ToResponse - конвертация User в UserResponse
func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:                 u.ID,
		Username:           u.Username,
		Email:              u.Email,
		AvatarURL:          u.AvatarURL,
		Role:               u.Role,
		CreatedAt:          u.CreatedAt,
		BooksRead:          u.BooksRead,
		ReviewsCount:       u.ReviewsCount,
		LikesReceived:      u.LikesReceived,
		ProfileVisibility:  u.ProfileVisibility,
		ActivityVisibility: u.ActivityVisibility,
	}
}
