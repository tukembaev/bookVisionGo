package interfaces

import (
	"context"

	"github.com/tukembaev/bookVisionGo/internal/models"
)

// UserRepository - интерфейс для работы с пользователями
type UserRepository interface {
	// Create - создание нового пользователя
	Create(ctx context.Context, user *models.User) error

	// GetByID - получение пользователя по ID
	GetByID(ctx context.Context, id string) (*models.User, error)

	// GetByUsername - получение пользователя по username
	GetByUsername(ctx context.Context, username string) (*models.User, error)

	// Update - обновление данных пользователя
	Update(ctx context.Context, user *models.User) error

	// Delete - удаление пользователя
	Delete(ctx context.Context, id string) error

	// List - получение списка пользователей с пагинацией
	List(ctx context.Context, limit, offset int) ([]*models.User, error)

	// Count - подсчет общего количества пользователей
	Count(ctx context.Context) (int, error)

	// VerifyPassword - проверка пароля пользователя
	VerifyPassword(ctx context.Context, username, password string) (*models.User, error)
}
