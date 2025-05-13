package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"project/backend/db"
	"project/backend/handlers"
	"project/backend/middleware"
)

func main() {
	// Инициализация базы данных SQLite
	db.InitDB()

	// Настройка Gin
	r := gin.Default()

	// Настройка CORS
	r.Use(cors.New(cors.Config{
    	AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000", "https://source.unsplash.com"},
    	AllowMethods:     []string{"GET", "POST", "OPTIONS"},
    	AllowHeaders:     []string{"Origin", "Content-Type", "X-CSRF-Token", "Cookie"},
    	ExposeHeaders:    []string{"X-CSRF-Token", "Set-Cookie"},
    	AllowCredentials: true,
	}))

	// Публичные маршруты
	r.GET("/csrf-token", handlers.GetCSRFToken)
	r.POST("/register", middleware.CSRFMiddleware(), handlers.Register)
	r.POST("/login", middleware.CSRFMiddleware(), handlers.Login)
	r.POST("/logout", handlers.Logout)

	// Защищенные маршруты (требуется JWT)
	protected := r.Group("/").Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", handlers.GetProfile)
	}

	// Запуск сервера
	log.Println("Запускаем сервер на :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
