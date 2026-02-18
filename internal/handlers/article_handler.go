package handlers

import "github.com/tukembaev/bookVisionGo/internal/repositories/interfaces"

type ArticleHandler struct {
	articleHandler interfaces.ArticleRepository
}

func NewArticleHandler(articleHandler interfaces.ArticleRepository) *ArticleHandler {
	return &ArticleHandler{
		articleHandler: articleHandler,
	}
}
