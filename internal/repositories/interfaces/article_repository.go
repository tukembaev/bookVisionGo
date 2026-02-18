package interfaces

import (
	"context"

	"github.com/tukembaev/bookVisionGo/internal/models"
)

type ArticleRepository interface {
	GetList(ctx context.Context) ([]*models.Article, error)
	CreateArticle(ctx context.Context, article *models.Article) error
	GetByID(ctx context.Context, id string) (*models.Article, error)
	Update(ctx context.Context, article *models.Article) error
	Delete(ctx context.Context, id string) error
}
