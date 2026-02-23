package repositories

import (
	"context"
	"fmt"
	"strconv"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tukembaev/bookVisionGo/internal/models"
	"github.com/tukembaev/bookVisionGo/internal/repositories/interfaces"
)

type ArticleRepository struct {
	pool *pgxpool.Pool
}

func NewArticleRepository(pool *pgxpool.Pool) interfaces.ArticleRepository {
	return &ArticleRepository{
		pool: pool,
	}
}

// GetList - получение списка статей
func (r *ArticleRepository) GetList(ctx context.Context, sortBy, order string, limit string) ([]*models.ArticleListItem, error) {
	// 1. Белый список полей для сортировки (Защита от SQL-инъекций)
	allowedColumns := map[string]string{
		"likes":      "likes",
		"views":      "views",
		"created_at": "created_at",
		"newest":     "created_at", // алиас для удобства
	}

	column, ok := allowedColumns[sortBy]
	if !ok {
		column = "created_at" // Сортировка по умолчанию
	}

	// 2. Валидация направления
	if order != "asc" {
		order = "desc" // По умолчанию самые свежие/популярные сверху
	}

	// 3. Валидация и конвертация limit
	limitInt := 10 // по умолчанию
	if limit != "" {
		if parsedLimit, err := strconv.Atoi(limit); err == nil && parsedLimit > 0 {
			limitInt = parsedLimit
		}
	}

	query := fmt.Sprintf(`SELECT id, title, type, author_id, book_id, excerpt, likes, views, cover_url 
							FROM articles 
							ORDER BY %s %s 
							LIMIT %d`, column, order, limitInt)

	var articles []*models.ArticleListItem
	err := pgxscan.Select(ctx, r.pool, &articles, query)
	if err != nil {
		return nil, fmt.Errorf("failed to select articles: %w", err)
	}
	return articles, nil
}

// GetByID - получение статьи по ID
func (r *ArticleRepository) GetByID(ctx context.Context, id string) (*models.Article, error) {
	query := `SELECT id, title, type, author_id, book_id, excerpt, created_at, likes, views, reading_minutes, cover_url, verified, verification_type, no_spoilers, readiness, content 
				FROM articles 
			    WHERE id = $1 `
	var article models.Article
	err := pgxscan.Get(ctx, r.pool, &article, query, id)

	if err != nil {
		if pgxscan.NotFound(err) {
			return nil, fmt.Errorf("article not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get article: %w", err)
	}
	return &article, nil
}

// CreateArticle - создание новой статьи
func (r *ArticleRepository) CreateArticle(ctx context.Context, article *models.Article) error {
	// TODO: Реализовать создание статьи
	return fmt.Errorf("CreateArticle not implemented yet")
}

// Update - обновление статьи
func (r *ArticleRepository) Update(ctx context.Context, article *models.Article) error {
	// TODO: Реализовать обновление статьи
	return fmt.Errorf("Update not implemented yet")
}

// Delete - удаление статьи
func (r *ArticleRepository) Delete(ctx context.Context, id string) error {
	// TODO: Реализовать удаление статьи
	return fmt.Errorf("Delete not implemented yet")
}
