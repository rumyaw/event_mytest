package handlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"log"
	"net/http"
	"project/backend/db"
	"project/backend/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("your-secret-key") // Замени на безопасный ключ

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Invalid input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Проверка на пустые поля
	if user.Name == "" || user.Surname == "" || user.Email == "" || user.Password == "" || user.Role == "" {
		log.Printf("Missing required fields: %+v", user)
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	// Проверка корректности роли
	validRoles := map[string]bool{"participant": true, "organizer": true, "moderator": true}
	if !validRoles[user.Role] {
		log.Printf("Invalid role: %s", user.Role)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
		return
	}

	// Генерация соли
	saltBytes := make([]byte, 16)
	_, err := rand.Read(saltBytes)
	if err != nil {
		log.Printf("Failed to generate salt: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate salt"})
		return
	}
	salt := string(saltBytes)

	// Хэширование пароля
	hashedPassword := db.HashPassword(user.Password, salt)

	// Сохранение пользователя
	_, err = db.DB.Exec("INSERT INTO users (name, surname, email, password, role, salt) VALUES (?, ?, ?, ?, ?, ?)",
		user.Name, user.Surname, user.Email, hashedPassword, user.Role, salt)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.email" {
			log.Printf("Duplicate email: %s", user.Email)
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		} else {
			log.Printf("Failed to register user: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
		}
		return
	}

	log.Printf("User registered successfully: %s", user.Email)
	c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Invalid input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	err := db.DB.QueryRow("SELECT id, name, surname, email, password, role, salt FROM users WHERE email = ?", input.Email).
		Scan(&user.ID, &user.Name, &user.Surname, &user.Email, &user.Password, &user.Role, &user.Salt)
	if err == sql.ErrNoRows {
		log.Printf("User not found: %s", input.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	} else if err != nil {
		log.Printf("Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	hashedInput := db.HashPassword(input.Password, user.Salt)
	if hashedInput != user.Password {
		log.Printf("Invalid password for email: %s", input.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Printf("Failed to generate token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie("token", tokenString, 3600*24, "/", "", false, true)
	log.Printf("User logged in: %s", user.Email)
	c.JSON(http.StatusOK, gin.H{"message": "Logged in", "role": user.Role})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	log.Println("User logged out")
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}

func GetCSRFToken(c *gin.Context) {
    log.Println("Получен запрос на /csrf-token")

    csrfToken, err := c.Cookie("csrf_token")
    if err != nil || csrfToken == "" {
        tokenBytes := make([]byte, 32)
        _, err := rand.Read(tokenBytes)
        if err != nil {
            log.Printf("Failed to generate CSRF token: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate CSRF token"})
            return
        }
        csrfToken = base64.StdEncoding.EncodeToString(tokenBytes)

        // Устанавливаем куку без указания домена
        c.SetCookie("csrf_token", csrfToken, 3600, "/", "", false, false)
        log.Printf("Установлен новый CSRF-токен: %s", csrfToken)
    } else {
        log.Printf("Используется существующий CSRF-токен: %s", csrfToken)
    }

    c.Header("X-CSRF-Token", csrfToken)
    c.JSON(http.StatusOK, gin.H{"csrf_token": csrfToken})
}
