package middleware

import (

	"net/http"
	"log"
	"github.com/gin-gonic/gin"
)

func CSRFMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        if c.Request.Method == "OPTIONS" {
            c.Next()
            return
        }

        if c.Request.Method == "POST" {
            headerToken := c.GetHeader("X-CSRF-Token")
            if headerToken == "" {
                c.JSON(http.StatusForbidden, gin.H{"error": "Missing CSRF token in header"})
                c.Abort()
                return
            }
            // Временно пропускаем проверку куки
            log.Printf("Получен X-CSRF-Token: %s", headerToken)
        }

        c.Next()
    }
}
