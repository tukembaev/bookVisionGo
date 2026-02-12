package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tukembaev/bookVisionGo/internal/config"
)

// Database - структура для работы с базой данных
type Database struct {
	Pool *pgxpool.Pool
}

// NewDatabase - создание нового подключения к базе данных
func NewDatabase(cfg *config.Config) (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Используем DatabaseURL() из конфига
	dbURL := cfg.DatabaseURL()

	log.Printf("Connecting to database: %s", dbURL)

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	// Проверка подключения
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	log.Printf("Successfully connected to database")

	return &Database{
		Pool: pool,
	}, nil
}

// Close - закрытие подключения к базе данных
func (d *Database) Close() {
	if d.Pool != nil {
		d.Pool.Close()
		log.Println("Database connection closed")
	}
}

// GetPool - получение connection pool
func (d *Database) GetPool() *pgxpool.Pool {
	return d.Pool
}

// HealthCheck - проверка здоровья базы данных
func (d *Database) HealthCheck(ctx context.Context) error {
	if d.Pool == nil {
		return fmt.Errorf("database pool is nil")
	}

	return d.Pool.Ping(ctx)
}

// GetStats - получение статистики подключения
func (d *Database) GetStats() *pgxpool.Stat {
	return d.Pool.Stat()
}
