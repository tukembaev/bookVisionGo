package db

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tukembaev/bookVisionGo/internal/models"
)

// SeedBooks - заполнение базы данных начальными книгами
func SeedBooks(ctx context.Context, pool *pgxpool.Pool) error {
	books := []models.Book{
		{
			ID:               uuid.New().String(),
			Title:            "Отцы и дети",
			OriginalTitle:    stringPtr("Fathers and Sons"),
			Author:           "Иван Тургенев",
			Year:             intPtr(1862),
			Genres:           []string{"Роман", "Драма", "Философия"},
			AgeRating:        ageRatingPtr(models.AgeRating12),
			AuthorCountry:    stringPtr("Россия"),
			Description:      "Роман о конфликте поколений и столкновении идей в России середины XIX века.",
			CoverURL:         stringPtr("https://cdn.azbooka.ru/cv/w1100/61a7ee1f-7c15-412b-b3b0-37aec56fabc2.jpg"),
			PagesCount:       320,
			Verified:         true,
			VerificationType: verificationTypePtr(models.VerificationTypeAI),
			CreatedBy:        nil,
			CreatedAt:        time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			Tags:             []string{"философия", "общество", "конфликт"},
			AverageRating:    8.6,
			RatingCount:      1240,
		},
		{
			ID:               uuid.New().String(),
			Title:            "Преступление и наказание",
			Author:           "Фёдор Достоевский",
			Year:             intPtr(1866),
			Genres:           []string{"Роман", "Психология", "Философия"},
			AgeRating:        ageRatingPtr(models.AgeRating16),
			AuthorCountry:    stringPtr("Россия"),
			Description:      "История внутреннего кризиса и морального выбора, разворачивающаяся вокруг преступления.",
			CoverURL:         stringPtr("https://flibusta.su/b/img/big/208394.jpg"),
			PagesCount:       560,
			Verified:         true,
			VerificationType: verificationTypePtr(models.VerificationTypeAI),
			CreatedBy:        nil,
			CreatedAt:        time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC),
			Tags:             []string{"психология", "вина", "искупление"},
			AverageRating:    9.1,
			RatingCount:      2305,
		},
		{
			ID:               uuid.New().String(),
			Title:            "Мастер и Маргарита",
			Author:           "Михаил Булгаков",
			Year:             intPtr(1967),
			Genres:           []string{"Роман", "Фантастика", "Философия"},
			AgeRating:        ageRatingPtr(models.AgeRating16),
			AuthorCountry:    stringPtr("Россия"),
			Description:      "Сатира и мистический роман, переплетающий несколько линий и смысловых пластов.",
			CoverURL:         stringPtr("https://cdn.azbooka.ru/cv/w1100/98fa6b42-e86d-4f17-9376-25e98cc784e5.jpg"),
			PagesCount:       410,
			Verified:         true,
			VerificationType: verificationTypePtr(models.VerificationTypeCommunity),
			CreatedBy:        nil,
			CreatedAt:        time.Date(2026, 1, 3, 0, 0, 0, 0, time.UTC),
			Tags:             []string{"мистика", "сатира", "любовь"},
			AverageRating:    9.0,
			RatingCount:      3102,
		},
		{
			ID:               uuid.New().String(),
			Title:            "Отверженные",
			OriginalTitle:    stringPtr("Les Misérables"),
			Author:           "Виктор Гюго",
			Year:             intPtr(1862),
			Genres:           []string{"Роман", "Драма", "Исторический"},
			AgeRating:        ageRatingPtr(models.AgeRating12),
			AuthorCountry:    stringPtr("Франция"),
			Description:      "Эпический роман о милосердии, справедливости и судьбах людей на фоне эпохи.",
			CoverURL:         stringPtr("https://cdn.azbooka.ru/cv/w1100/8e4f70bd-f412-4f3c-a9cf-b1755601fb97.jpg"),
			PagesCount:       1240,
			Verified:         true,
			VerificationType: verificationTypePtr(models.VerificationTypeAI),
			CreatedBy:        nil,
			CreatedAt:        time.Date(2026, 1, 4, 0, 0, 0, 0, time.UTC),
			Tags:             []string{"эпос", "общество", "история"},
			AverageRating:    8.9,
			RatingCount:      980,
		},
		{
			ID:               uuid.New().String(),
			Title:            "Норвежский лес",
			OriginalTitle:    stringPtr("Norwegian Wood"),
			Author:           "Харуки Мураками",
			Year:             intPtr(1987),
			Genres:           []string{"Роман", "Драма"},
			AgeRating:        ageRatingPtr(models.AgeRating16),
			AuthorCountry:    stringPtr("Япония"),
			Description:      "Тихий роман о взрослении, памяти и потерях.",
			CoverURL:         stringPtr("https://flibusta.su/b/img/big/589.jpg"),
			PagesCount:       384,
			Verified:         true,
			VerificationType: verificationTypePtr(models.VerificationTypeCommunity),
			CreatedBy:        nil,
			CreatedAt:        time.Date(2026, 1, 5, 0, 0, 0, 0, time.UTC),
			Tags:             []string{"взросление", "память"},
			AverageRating:    8.1,
			RatingCount:      1405,
		},
	}

	for _, book := range books {
		err := insertBook(ctx, pool, &book)
		if err != nil {
			return fmt.Errorf("ошибка вставки книги %s: %w", book.Title, err)
		}
		fmt.Printf("Книга '%s' успешно добавлена\n", book.Title)
	}

	return nil
}

func insertBook(ctx context.Context, pool *pgxpool.Pool, book *models.Book) error {
	query := `
		INSERT INTO books (
			id, title, original_title, author, year, genres, age_rating,
			author_country, description, cover_url, pages_count, tags,
			verified, verification_type, created_by, created_at,
			average_rating, rating_count
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18
		) ON CONFLICT (id) DO NOTHING`

	_, err := pool.Exec(ctx, query,
		book.ID,
		book.Title,
		book.OriginalTitle,
		book.Author,
		book.Year,
		book.Genres,
		book.AgeRating,
		book.AuthorCountry,
		book.Description,
		book.CoverURL,
		book.PagesCount,
		book.Tags,
		book.Verified,
		book.VerificationType,
		book.CreatedBy,
		book.CreatedAt,
		book.AverageRating,
		book.RatingCount,
	)

	return err
}

// Вспомогательные функции для создания указателей
func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func ageRatingPtr(ar models.AgeRating) *models.AgeRating {
	return &ar
}

func verificationTypePtr(vt models.VerificationType) *models.VerificationType {
	return &vt
}
