package repositories

import (
	"context"
	"fmt"

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
func (r *ArticleRepository) GetList(ctx context.Context) ([]*models.Article, error) {
	// TODO: Реализовать получение списка статей
	query := `SELECT id , title , type , author_id , excerpt , created_at , likes , views ,rating , cover_url FROM articles`
	var articles []*models.Article
	err := pgxscan.Select(ctx, r.pool, &articles, query)
	if err != nil {
		return nil, fmt.Errorf("failed to select articles: %w", err)
	}
	return articles, nil
}

// CreateArticle - создание новой статьи
func (r *ArticleRepository) CreateArticle(ctx context.Context, article *models.Article) error {
	// TODO: Реализовать создание статьи
	return fmt.Errorf("CreateArticle not implemented yet")
}

// GetByID - получение статьи по ID
func (r *ArticleRepository) GetByID(ctx context.Context, id string) (*models.Article, error) {
	// TODO: Реализовать получение статьи
	return nil, fmt.Errorf("GetByID not implemented yet")
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
