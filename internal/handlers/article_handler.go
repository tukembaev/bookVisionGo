package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tukembaev/bookVisionGo/internal/repositories/interfaces"
)

type ArticleHandler struct {
	articleHandler interfaces.ArticleRepository
}

func NewArticleHandler(articleHandler interfaces.ArticleRepository) *ArticleHandler {
	return &ArticleHandler{
		articleHandler: articleHandler,
	}
}

// GetArticles - получение списка статей
// @Summary Получение списка статей
// @Description Получение списка статей
// @Tags articles
// @Accept json
// @Produce json
// @Param sort query string false "Поле для сортировки (views, likes, created_at)"
// @Param order query string false "Порядок сортировки (asc, desc)"
// @Param limit query int false "Количество статей"
// @Success 200 {object} []models.ArticleListItem
// @Failure 400 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/articles [get]
func (h *ArticleHandler) GetArticles(c *gin.Context) {
	sortBy := c.Query("sort")
	limit := c.Query("limit")
	order := c.Query("order")
	articles, err := h.articleHandler.GetList(c.Request.Context(), sortBy, order, limit)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, articles)
}

// GetArticleById - получение статьи по ID
// @Summary Получение статьи по ID
// @Description Получение статьи по ID f80b90a5-a9e3-4347-9d0c-1a8b0abdfbf2
// @Tags articles
// @Accept json
// @Produce json
// @Param id path string true "ID статьи"
// @Success 200 {object} models.Article
// @Failure 400 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/articles/{id} [get]
func (h *ArticleHandler) GetArticleById(c *gin.Context) {
	article, err := h.articleHandler.GetByID(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, article)
}
