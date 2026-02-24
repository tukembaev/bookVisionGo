package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tukembaev/bookVisionGo/internal/models"
	"github.com/tukembaev/bookVisionGo/internal/repositories/interfaces"
)

// BookRepository - реализация репозитория для книг
type BookRepository struct {
	pool *pgxpool.Pool
}

// NewBookRepository - создание нового BookRepository
func NewBookRepository(pool *pgxpool.Pool) interfaces.BookRepository {
	return &BookRepository{
		pool: pool,
	}
}

// Create - создание новой книги
func (r *BookRepository) Create(ctx context.Context, book *models.Book) error {
	query := `
		INSERT INTO books (
			title, original_title, author, year, genres, age_rating,
			author_country, description, cover_url, pages_count, tags,
			verified, verification_type, created_by, created_at,
			average_rating, rating_count
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17
		) RETURNING id`

	var id string
	err := r.pool.QueryRow(ctx, query,
		book.Title, book.OriginalTitle, book.Author, book.Year,
		book.Genres, book.AgeRating, book.AuthorCountry, book.Description,
		book.CoverURL, book.PagesCount, book.Tags, book.Verified,
		book.VerificationType, book.CreatedBy, book.CreatedAt,
		book.AverageRating, book.RatingCount,
	).Scan(&id)

	if err != nil {
		return fmt.Errorf("failed to create book: %w", err)
	}

	book.ID = id
	return nil
}

// GetByID - получение книги по ID
func (r *BookRepository) GetByID(ctx context.Context, id string) (*models.Book, error) {
	query := `
		SELECT id, title, original_title, author, year, genres, age_rating,
			   author_country, description, cover_url, pages_count, tags,
			   verified, verification_type, created_by, created_at,
			   average_rating, rating_count
		FROM books 
		WHERE id = $1`

	var book models.Book
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&book.ID, &book.Title, &book.OriginalTitle, &book.Author, &book.Year,
		&book.Genres, &book.AgeRating, &book.AuthorCountry, &book.Description,
		&book.CoverURL, &book.PagesCount, &book.Tags, &book.Verified,
		&book.VerificationType, &book.CreatedBy, &book.CreatedAt,
		&book.AverageRating, &book.RatingCount,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get book by id: %w", err)
	}

	return &book, nil
}

// Update - обновление книги
func (r *BookRepository) Update(ctx context.Context, book *models.Book) error {
	query := `
        UPDATE books SET 
            title = $2, original_title = $3, author = $4, year = $5, genres = $6, 
            age_rating = $7, author_country = $8, description = $9, cover_url = $10, 
            pages_count = $11, tags = $12, verified = $13, verification_type = $14,
            average_rating = $15, rating_count = $16
        WHERE id = $1`

	cmdTag, err := r.pool.Exec(ctx, query,
		book.ID, book.Title, book.OriginalTitle, book.Author, book.Year,
		book.Genres, book.AgeRating, book.AuthorCountry, book.Description,
		book.CoverURL, book.PagesCount, book.Tags, book.Verified,
		book.VerificationType, book.AverageRating, book.RatingCount,
	)

	if err != nil {
		return fmt.Errorf("failed to update book: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("book with id %s not found", book.ID)
	}

	return nil
}

// Delete - удаление книги
func (r *BookRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM books WHERE id = $1`

	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete book: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("book with id %s not found", id)
	}

	return nil
}

// List - получение списка книг с фильтрацией и пагинацией
func (r *BookRepository) List(ctx context.Context, filters interfaces.BookFilters, limit, offset int) ([]*models.Book, error) {
	// Реальный SQL запрос к БД
	query := `
		SELECT id, title, original_title, author, year, genres, age_rating,
			   author_country, description, cover_url, pages_count, tags,
			   verified, verification_type, created_by, created_at,
			   average_rating, rating_count
		FROM books 
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`

	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query books: %w", err)
	}
	defer rows.Close()

	var books []*models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(
			&book.ID, &book.Title, &book.OriginalTitle, &book.Author, &book.Year,
			&book.Genres, &book.AgeRating, &book.AuthorCountry, &book.Description,
			&book.CoverURL, &book.PagesCount, &book.Tags, &book.Verified,
			&book.VerificationType, &book.CreatedBy, &book.CreatedAt,
			&book.AverageRating, &book.RatingCount,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan book: %w", err)
		}
		books = append(books, &book)
	}

	return books, nil
}

// Count - подсчет количества книг с фильтрацией
func (r *BookRepository) Count(ctx context.Context, filters interfaces.BookFilters) (int, error) {
	// Реальный SQL запрос для подсчета
	query := `SELECT COUNT(*) FROM books`

	var count int
	err := r.pool.QueryRow(ctx, query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count books: %w", err)
	}

	return count, nil
}

// GetParts - получение частей книги
func (r *BookRepository) GetParts(ctx context.Context, bookID string) ([]*models.BookPart, error) {
	query := `
		SELECT id, book_id, title, content, order_num, page_start, page_end, mood_tags, average_rating
		FROM book_parts 
		WHERE book_id = $1
		ORDER BY order_num`

	rows, err := r.pool.Query(ctx, query, bookID)
	if err != nil {
		return nil, fmt.Errorf("failed to query book parts: %w", err)
	}
	defer rows.Close()

	var parts []*models.BookPart
	for rows.Next() {
		var part models.BookPart
		err := rows.Scan(
			&part.ID, &part.BookID, &part.Title, &part.Content, &part.OrderNum,
			&part.PageStart, &part.PageEnd, &part.MoodTags, &part.AverageRating,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan book part: %w", err)
		}
		parts = append(parts, &part)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating book parts: %w", err)
	}

	return parts, nil
}

// GetPartByID - получение части книги по ID
func (r *BookRepository) GetPartByID(ctx context.Context, partID string) (*models.BookPart, error) {
	// TODO: Реализовать получение части книги
	return nil, fmt.Errorf("GetPartByID not implemented yet")
}

// CreatePart - создание новой части книги
func (r *BookRepository) CreatePart(ctx context.Context, part *models.BookPart) error {
	// TODO: Реализовать создание части книги
	return fmt.Errorf("CreatePart not implemented yet")
}

// UpdatePart - обновление части книги
func (r *BookRepository) UpdatePart(ctx context.Context, part *models.BookPart) error {
	// TODO: Реализовать обновление части книги
	return fmt.Errorf("UpdatePart not implemented yet")
}

// DeletePart - удаление части книги
func (r *BookRepository) DeletePart(ctx context.Context, partID string) error {
	// TODO: Реализовать удаление части книги
	return fmt.Errorf("DeletePart not implemented yet")
}
