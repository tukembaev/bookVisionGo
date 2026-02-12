package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tukembaev/bookVisionGo/internal/models"
	"github.com/tukembaev/bookVisionGo/internal/repositories/interfaces"
)

// BookHandler - обработчики для работы с книгами
type BookHandler struct {
	bookRepo interfaces.BookRepository
}

// NewBookHandler - создание нового BookHandler
func NewBookHandler(bookRepo interfaces.BookRepository) *BookHandler {
	return &BookHandler{
		bookRepo: bookRepo,
	}
}

// CreateBook - создание новой книги (требует прав moderator/admin)
// @Summary Создание книги
// @Description Создание новой книги в каталоге
// @Tags books
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Param request body models.CreateBookRequest true "Данные книги"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Router /api/books [post]
func (h *BookHandler) CreateBook(c *gin.Context) {
	var req models.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Создание книги
	book := &models.Book{
		Title:         req.Title,
		OriginalTitle: req.OriginalTitle,
		Author:        req.Author,
		Year:          req.Year,
		Genres:        req.Genres,
		AgeRating:     req.AgeRating,
		AuthorCountry: req.AuthorCountry,
		Description:   req.Description,
		CoverURL:      req.CoverURL,
		PagesCount:    req.PagesCount,
		Tags:          req.Tags,
		Verified:      false, // По умолчанию не верифицирована
	}

	err := h.bookRepo.Create(c.Request.Context(), book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"book": book.ToResponse(),
	})
}

// GetBooks - получение списка книг с фильтрацией и пагинацией
// @Summary Получение списка книг
// @Description Получение каталога книг с фильтрацией
// @Tags books
// @Produce json
// @Param genre query string false "Фильтр по жанру"
// @Param author query string false "Фильтр по автору"
// @Param year query int false "Фильтр по году"
// @Param min_rating query number false "Минимальный рейтинг"
// @Param verified query bool false "Только верифицированные"
// @Param search query string false "Поиск по названию"
// @Param limit query int false "Лимит" default(20)
// @Param offset query int false "Смещение" default(0)
// @Success 200 {object} map[string]interface{}
// @Router /api/books [get]
func (h *BookHandler) GetBooks(c *gin.Context) {
	fmt.Println("GetBooks called!") // Debug лог

	// Парсинг query параметров
	filters := interfaces.BookFilters{}

	if genre := c.Query("genre"); genre != "" {
		filters.Genre = &genre
	}
	if author := c.Query("author"); author != "" {
		filters.Author = &author
	}
	if yearStr := c.Query("year"); yearStr != "" {
		if year, err := strconv.Atoi(yearStr); err == nil {
			filters.Year = &year
		}
	}
	if minRatingStr := c.Query("min_rating"); minRatingStr != "" {
		if minRating, err := strconv.ParseFloat(minRatingStr, 64); err == nil {
			filters.MinRating = &minRating
		}
	}
	if verifiedStr := c.Query("verified"); verifiedStr != "" {
		if verified, err := strconv.ParseBool(verifiedStr); err == nil {
			filters.Verified = &verified
		}
	}
	if search := c.Query("search"); search != "" {
		filters.Search = &search
	}

	// Пагинация
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	// Получение книг из БД
	books, err := h.bookRepo.List(c.Request.Context(), filters, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Конвертация в response
	bookResponses := make([]*models.BookResponse, len(books))
	for i, book := range books {
		bookResponses[i] = book.ToResponse()
	}

	// Получение общего количества
	total, err := h.bookRepo.Count(c.Request.Context(), filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"books":  bookResponses,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	})
}

// GetBook - получение книги по ID
// @Summary Получение книги
// @Description Получение детальной информации о книге
// @Tags books
// @Produce json
// @Param id path string true "ID книги"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/books/{id} [get]
func (h *BookHandler) GetBook(c *gin.Context) {
	id := c.Param("id")

	book, err := h.bookRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"book": book.ToResponse(),
	})
}

// UpdateBook - обновление книги (требует прав moderator/admin)
// @Summary Обновление книги
// @Description Обновление информации о книге
// @Tags books
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Param id path string true "ID книги"
// @Param request body models.UpdateBookRequest true "Данные для обновления"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/books/{id} [put]
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var req models.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Получение существующей книги
	book, err := h.bookRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Обновление полей если они указаны
	if req.Title != nil {
		book.Title = *req.Title
	}
	if req.OriginalTitle != nil {
		book.OriginalTitle = req.OriginalTitle
	}
	if req.Author != nil {
		book.Author = *req.Author
	}
	if req.Year != nil {
		book.Year = req.Year
	}
	if req.Genres != nil {
		book.Genres = req.Genres
	}
	if req.AgeRating != nil {
		book.AgeRating = req.AgeRating
	}
	if req.AuthorCountry != nil {
		book.AuthorCountry = req.AuthorCountry
	}
	if req.Description != nil {
		book.Description = *req.Description
	}
	if req.CoverURL != nil {
		book.CoverURL = req.CoverURL
	}
	if req.PagesCount != nil {
		book.PagesCount = *req.PagesCount
	}
	if req.Tags != nil {
		book.Tags = req.Tags
	}
	if req.Verified != nil {
		book.Verified = *req.Verified
	}
	if req.VerificationType != nil {
		book.VerificationType = req.VerificationType
	}

	err = h.bookRepo.Update(c.Request.Context(), book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"book": book.ToResponse(),
	})
}

// DeleteBook - удаление книги (требует прав admin)
// @Summary Удаление книги
// @Description Удаление книги из каталога
// @Tags books
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Param id path string true "ID книги"
// @Success 200 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/books/{id} [delete]
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	err := h.bookRepo.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})
}

// GetBookParts - получение частей книги
// @Summary Получение частей книги
// @Description Получение списка глав/частей книги
// @Tags books
// @Produce json
// @Param id path string true "ID книги"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/books/{id}/parts [get]
func (h *BookHandler) GetBookParts(c *gin.Context) {
	bookID := c.Param("id")

	parts, err := h.bookRepo.GetParts(c.Request.Context(), bookID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Конвертация в response
	partResponses := make([]*models.BookPartResponse, len(parts))
	for i, part := range parts {
		partResponses[i] = part.ToResponse()
	}

	c.JSON(http.StatusOK, gin.H{
		"parts": partResponses,
	})
}

// GetBookPart - получение части книги по ID
// @Summary Получение части книги
// @Description Получение информации о конкретной главе/части
// @Tags books
// @Produce json
// @Param id path string true "ID книги"
// @Param partId path string true "ID части"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/books/{id}/parts/{partId} [get]
func (h *BookHandler) GetBookPart(c *gin.Context) {
	partID := c.Param("partId")

	part, err := h.bookRepo.GetPartByID(c.Request.Context(), partID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Part not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"part": part.ToResponse(),
	})
}
