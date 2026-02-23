package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/tukembaev/bookVisionGo/docs"
	"github.com/tukembaev/bookVisionGo/internal/config"
	"github.com/tukembaev/bookVisionGo/internal/db"
	"github.com/tukembaev/bookVisionGo/internal/handlers"
	"github.com/tukembaev/bookVisionGo/internal/repositories"
	"github.com/tukembaev/bookVisionGo/internal/services"
	"github.com/tukembaev/bookVisionGo/internal/utils"
	"github.com/tukembaev/bookVisionGo/pkg/api"
)

// @title Book Vision Go API
// @version 1.0
// @description API документация для платформы Book Vision - социальная платформа для чтения книг
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @schemes http https
// @securityDefinitions.http BearerAuth
// @scheme bearer
// @bearerFormat JWT

func main() {
	// Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Вывод информации о хосте и конфигурации
	fmt.Printf("📋 Configuration loaded:\n")
	fmt.Printf("   - Server Port: '%s'\n", cfg.Server.Port)
	fmt.Printf("   - Server Mode: '%s'\n", cfg.Server.Mode)
	fmt.Printf("   - DB Host: '%s'\n", cfg.Database.Host)
	fmt.Printf("   - DB Port: '%s'\n", cfg.Database.Port)

	// Если порт пустой, используем дефолтный
	port := cfg.Server.Port
	if port == "" {
		port = "8080"
		fmt.Printf("⚠️  Port not configured, using default: %s\n", port)
	}

	fmt.Printf("🚀 Server will start on: http://localhost:%s\n", port)
	fmt.Printf("📚 Swagger UI will be available at: http://localhost:%s/swagger/index.html\n", port)
	fmt.Printf("🔍 Health check: http://localhost:%s/health\n", port)

	// Настройка Gin режима
	gin.SetMode(cfg.Server.Mode)

	// Подключение к базе данных
	database, err := db.NewDatabase(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	// Инициализация зависимостей
	jwtUtils := utils.NewJWTUtils(cfg)

	// Репозитории
	userRepo := repositories.NewUserRepository(database.GetPool())
	bookRepo := repositories.NewBookRepository(database.GetPool()) // Включаем BookRepository
	articleRepo := repositories.NewArticleRepository(database.GetPool())
	// Сервисы
	authService := services.NewAuthService(userRepo, jwtUtils)

	// Handlers
	authHandler := handlers.NewAuthHandler(authService)
	bookHandler := handlers.NewBookHandler(bookRepo)          // Настоящий handler с репозиторием
	articleHandler := handlers.NewArticleHandler(articleRepo) // Настоящий handler с репозиторием

	// Debug: проверим что handler не nil
	if bookHandler == nil {
		log.Fatal("bookHandler is nil!")
	}

	// Настройка роутов
	r := gin.Default()

	// Настройка CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:3000", // React dev server
		"http://localhost:5173", // Vite dev server
		"http://localhost:8080", // Swagger UI
	}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.AllowCredentials = true

	r.Use(cors.New(config))

	// Swagger документация
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api.SetupRoutes(r, authHandler, bookHandler, articleHandler, authService)

	// Запуск сервера
	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
