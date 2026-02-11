package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/tukembaev/bookVisionGo/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки конфигурации:", err)
	}

	dbURL := cfg.DatabaseURL()
	fmt.Println("Connection URL:", dbURL)

	cmd := exec.Command("migrate",
		"-path", "internal/db/migrations",
		"-database", dbURL,
		"up")

	fmt.Println("Running command:", cmd.String())

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Command output:", string(output))
		log.Fatal("Migration failed:", err)
	}

	fmt.Println("Migration successful!")
	fmt.Println("Output:", string(output))
}
