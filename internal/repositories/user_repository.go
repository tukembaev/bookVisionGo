package repositories

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"

	"github.com/tukembaev/bookVisionGo/internal/models"
	"github.com/tukembaev/bookVisionGo/internal/repositories/interfaces"
)

// userRepository - реализация UserRepository
type userRepository struct {
	db *pgxpool.Pool
}

// NewUserRepository - создание нового UserRepository
func NewUserRepository(db *pgxpool.Pool) interfaces.UserRepository {
	return &userRepository{
		db: db,
	}
}

// Create - создание нового пользователя
func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (id, username, email, password_hash, avatar_url, role, created_at, 
		                  books_read, reviews_count, likes_received, profile_visibility, activity_visibility)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`

	// Генерация UUID если не задан
	if user.ID == "" {
		user.ID = uuid.New().String()
	}

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Установка значений по умолчанию
	if user.Role == "" {
		user.Role = models.UserRoleUser
	}
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}

	_, err = r.db.Exec(ctx, query,
		user.ID,
		user.Username,
		user.Email,
		string(hashedPassword),
		user.AvatarURL,
		user.Role,
		user.CreatedAt,
		user.BooksRead,
		user.ReviewsCount,
		user.LikesReceived,
		user.ProfileVisibility,
		user.ActivityVisibility,
	)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	log.Printf("User created successfully: %s", user.Username)
	return nil
}

// GetByID - получение пользователя по ID
func (r *userRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	query := `
		SELECT id, username, email, password_hash, avatar_url, role, created_at,
			   books_read, reviews_count, likes_received, profile_visibility, activity_visibility
		FROM users 
		WHERE id = $1
	`

	var user models.User
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.AvatarURL,
		&user.Role,
		&user.CreatedAt,
		&user.BooksRead,
		&user.ReviewsCount,
		&user.LikesReceived,
		&user.ProfileVisibility,
		&user.ActivityVisibility,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

// GetByUsername - получение пользователя по username
func (r *userRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `
		SELECT id, username, email, password_hash, avatar_url, role, created_at,
			   books_read, reviews_count, likes_received, profile_visibility, activity_visibility
		FROM users 
		WHERE username = $1
	`

	var user models.User
	err := r.db.QueryRow(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.AvatarURL,
		&user.Role,
		&user.CreatedAt,
		&user.BooksRead,
		&user.ReviewsCount,
		&user.LikesReceived,
		&user.ProfileVisibility,
		&user.ActivityVisibility,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

// Update - обновление данных пользователя
func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	query := `
		UPDATE users 
		SET username = $2, avatar_url = $3, role = $4, 
		    books_read = $5, reviews_count = $6, likes_received = $7,
		    profile_visibility = $8, activity_visibility = $9
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query,
		user.ID,
		user.Username,
		user.AvatarURL,
		user.Role,
		user.BooksRead,
		user.ReviewsCount,
		user.LikesReceived,
		user.ProfileVisibility,
		user.ActivityVisibility,
	)

	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	log.Printf("User updated successfully: %s", user.Username)
	return nil
}

// Delete - удаление пользователя
func (r *userRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("user not found")
	}

	log.Printf("User deleted successfully: %s", id)
	return nil
}

// List - получение списка пользователей с пагинацией
func (r *userRepository) List(ctx context.Context, limit, offset int) ([]*models.User, error) {
	query := `
		SELECT id, username, email, password_hash, avatar_url, role, created_at,
			   books_read, reviews_count, likes_received, profile_visibility, activity_visibility
		FROM users 
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.PasswordHash,
			&user.AvatarURL,
			&user.Role,
			&user.CreatedAt,
			&user.BooksRead,
			&user.ReviewsCount,
			&user.LikesReceived,
			&user.ProfileVisibility,
			&user.ActivityVisibility,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, &user)
	}

	return users, nil
}

// Count - подсчет общего количества пользователей
func (r *userRepository) Count(ctx context.Context) (int, error) {
	query := `SELECT COUNT(*) FROM users`

	var count int
	err := r.db.QueryRow(ctx, query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count users: %w", err)
	}

	return count, nil
}

// VerifyPassword - проверка пароля пользователя
func (r *userRepository) VerifyPassword(ctx context.Context, username, password string) (*models.User, error) {
	user, err := r.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid password")
	}

	return user, nil
}
