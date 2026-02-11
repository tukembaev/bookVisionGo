// internal/config/config.go
package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DBConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port string `mapstructure:"SERVER_PORT"`
	Mode string `mapstructure:"GIN_MODE"`
}

type DBConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
	SSLMode  string `mapstructure:"DB_SSLMODE"`
}

type JWTConfig struct {
	SecretKey string `mapstructure:"JWT_SECRET"`
	ExpiresIn int    `mapstructure:"JWT_EXPIRES_IN"`
}

func Load() (*Config, error) {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")

	// Загрузка .env файла
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	} else {
		log.Printf("Successfully loaded .env file")
	}

	viper.AutomaticEnv()

	// Установка значений по умолчанию
	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("GIN_MODE", "debug")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_SSLMODE", "disable")
	viper.SetDefault("JWT_EXPIRES_IN", 24)

	// Отладка: выводим загруженные значения
	log.Printf("DB_HOST: %s", viper.GetString("DB_HOST"))
	log.Printf("DB_USER: %s", viper.GetString("DB_USER"))
	log.Printf("DB_PASSWORD: %s", viper.GetString("DB_PASSWORD"))
	log.Printf("DB_NAME: %s", viper.GetString("DB_NAME"))

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) DatabaseURL() string {
	// Используем прямой доступ к viper вместо полей структуры
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	name := viper.GetString("DB_NAME")
	sslMode := viper.GetString("DB_SSLMODE")

	// Отладка: выводим значения из структуры
	log.Printf("Config DB_HOST: %s", c.Database.Host)
	log.Printf("Config DB_USER: %s", c.Database.User)
	log.Printf("Config DB_PASSWORD: %s", c.Database.Password)
	log.Printf("Config DB_NAME: %s", c.Database.Name)
	log.Printf("Config DB_PORT: %s", c.Database.Port)
	log.Printf("Config DB_SSLMODE: %s", c.Database.SSLMode)

	// Отладка: выводим значения из viper
	log.Printf("Viper DB_HOST: %s", host)
	log.Printf("Viper DB_USER: %s", user)
	log.Printf("Viper DB_PASSWORD: %s", password)
	log.Printf("Viper DB_NAME: %s", name)

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user,
		password,
		host,
		port,
		name,
		sslMode,
	)
}
