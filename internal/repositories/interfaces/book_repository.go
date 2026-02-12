package interfaces

import (
	"context"
	"github.com/tukembaev/bookVisionGo/internal/models"
)

// BookRepository - интерфейс для работы с книгами
type BookRepository interface {
	// Create - создание новой книги
	Create(ctx context.Context, book *models.Book) error
	
	// GetByID - получение книги по ID
	GetByID(ctx context.Context, id string) (*models.Book, error)
	
	// Update - обновление данных книги
	Update(ctx context.Context, book *models.Book) error
	
	// Delete - удаление книги
	Delete(ctx context.Context, id string) error
	
	// List - получение списка книг с фильтрацией и пагинацией
	List(ctx context.Context, filters BookFilters, limit, offset int) ([]*models.Book, error)
	
	// Count - подсчет книг с фильтрацией
	Count(ctx context.Context, filters BookFilters) (int, error)
	
	// GetParts - получение частей книги
	GetParts(ctx context.Context, bookID string) ([]*models.BookPart, error)
	
	// GetPartByID - получение части книги по ID
	GetPartByID(ctx context.Context, partID string) (*models.BookPart, error)
	
	// CreatePart - создание новой части книги
	CreatePart(ctx context.Context, part *models.BookPart) error
	
	// UpdatePart - обновление части книги
	UpdatePart(ctx context.Context, part *models.BookPart) error
	
	// DeletePart - удаление части книги
	DeletePart(ctx context.Context, partID string) error
}

// BookFilters - фильтры для поиска книг
type BookFilters struct {
	Genre     *string
	Author    *string
	Year      *int
	MinRating *float64
	Verified  *bool
	Search    *string // для полнотекстового поиска
}
