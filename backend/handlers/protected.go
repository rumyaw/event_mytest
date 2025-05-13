package handlers

import (
    "net/http"
    "project/backend/db"
    "project/backend/models"
    "github.com/gin-gonic/gin"
    "strconv"
)

func GetProfile(c *gin.Context) {
    // достаём из контекста ID, который положил AuthMiddleware
    userIDraw, _ := c.Get("userID")
    // Gin отдаёт claims["id"] как float64, поэтому приводим
    userID, _ := strconv.Atoi(strconv.FormatFloat(userIDraw.(float64), 'f', -1, 64))

    // вытягиваем из БД остальные поля
    var user models.User
    err := db.DB.QueryRow(
        "SELECT id, name, surname, email, role FROM users WHERE id = ?",
        userID,
    ).Scan(&user.ID, &user.Name, &user.Surname, &user.Email, &user.Role)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить данные пользователя"})
        return
    }

    // возвращаем JSON без пароля и соли
    c.JSON(http.StatusOK, gin.H{
        "id":      user.ID,
        "name":    user.Name,
        "surname": user.Surname,
        "email":   user.Email,
        "role":    user.Role,
    })
}
