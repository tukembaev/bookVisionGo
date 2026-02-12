package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tukembaev/bookVisionGo/internal/handlers"
	"github.com/tukembaev/bookVisionGo/internal/middleware"
	"github.com/tukembaev/bookVisionGo/internal/models"
	"github.com/tukembaev/bookVisionGo/internal/services"
)

// SetupRoutes - настройка всех маршрутов API
func SetupRoutes(
	r *gin.Engine,
	authHandler *handlers.AuthHandler,
	bookHandler *handlers.BookHandler,
	authService *services.AuthService,
) {
	// Debug: проверим что handler не nil
	if bookHandler == nil {
		panic("bookHandler is nil in SetupRoutes!")
	}
	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API v1 group
	v1 := r.Group("/api")
	{
		// Auth routes (публичные)
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/refresh", authHandler.RefreshToken)

			// Требуют аутентификации
			authGroup := auth.Group("", middleware.AuthMiddleware(authService))
			{
				authGroup.GET("/profile", authHandler.GetProfile)
				authGroup.PUT("/profile", authHandler.UpdateProfile)
				authGroup.POST("/logout", authHandler.Logout)
			}
		}

		// Books routes
		books := v1.Group("/books")
		{
			// Сначала более конкретные маршруты
			books.GET("/:id/parts/:partId", bookHandler.GetBookPart)
			books.GET("/:id/parts", bookHandler.GetBookParts)

			// Затем общие маршруты
			books.GET("", bookHandler.GetBooks)
			books.GET("/:id", bookHandler.GetBook)

			// Защищенные маршруты
			booksGroup := books.Group("", middleware.AuthMiddleware(authService))
			{
				// Создание и обновление (требуют прав moderator+)
				moderatorGroup := booksGroup.Group("", middleware.RequireRole(models.UserRoleModerator))
				{
					moderatorGroup.POST("", bookHandler.CreateBook)
					moderatorGroup.PUT("/:id", bookHandler.UpdateBook)
				}

				// Удаление (требует прав admin)
				adminGroup := booksGroup.Group("", middleware.RequireRole(models.UserRoleAdmin))
				{
					adminGroup.DELETE("/:id", bookHandler.DeleteBook)
				}
			}
		}

		// Users routes (защищенные)
		users := v1.Group("/users", middleware.AuthMiddleware(authService))
		{
			users.GET("/me", func(c *gin.Context) {
				currentUser := middleware.GetCurrentUser(c)
				c.JSON(200, gin.H{
					"user": currentUser,
				})
			})

			// Admin только
			adminGroup := users.Group("", middleware.RequireRole(models.UserRoleAdmin))
			{
				adminGroup.GET("", func(c *gin.Context) {
					// Заглушка для получения списка пользователей
					c.JSON(200, gin.H{"message": "Admin users list"})
				})
			}
		}

		// Protected routes для будущих модулей
		protected := v1.Group("", middleware.AuthMiddleware(authService))
		{
			// Characters
			characters := protected.Group("/characters")
			{
				characters.GET("", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Characters list"})
				})
				characters.POST("", middleware.RequireRole(models.UserRoleModerator), func(c *gin.Context) {
					c.JSON(201, gin.H{"message": "Character created"})
				})
			}

			// Reviews
			reviews := protected.Group("/reviews")
			{
				reviews.GET("", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Reviews list"})
				})
				reviews.POST("", func(c *gin.Context) {
					c.JSON(201, gin.H{"message": "Review created"})
				})
			}

			// Articles
			articles := protected.Group("/articles")
			{
				articles.GET("", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Articles list"})
				})
				articles.POST("", middleware.RequireRole(models.UserRoleModerator), func(c *gin.Context) {
					c.JSON(201, gin.H{"message": "Article created"})
				})
			}

			// Challenges
			challenges := protected.Group("/challenges")
			{
				challenges.GET("", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Challenges list"})
				})
				challenges.POST("/:id/join", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Joined challenge"})
				})
			}

			// Playlists
			playlists := protected.Group("/playlists")
			{
				playlists.GET("", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Playlists list"})
				})
				playlists.POST("", func(c *gin.Context) {
					c.JSON(201, gin.H{"message": "Playlist created"})
				})
			}
		}
	}
}
