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

func (h *ArticleHandler) GetArticles(c *gin.Context) {
	articles, err := h.articleHandler.GetList(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, articles)
}
