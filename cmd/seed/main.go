package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tukembaev/bookVisionGo/internal/config"
	"github.com/tukembaev/bookVisionGo/internal/db"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Не удалось загрузить конфигурацию: %v", err)
	}

	// Подключение к базе данных
	database, err := db.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer database.Close()

	// Заполнение базы данных книгами
	err = db.SeedBooks(context.Background(), database.GetPool())
	if err != nil {
		log.Fatalf("Ошибка при заполнении базы данных: %v", err)
	}

	fmt.Println("База данных успешно заполнена!")
}
