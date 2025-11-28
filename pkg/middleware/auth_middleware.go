package middleware

import (
	"net/http"
	"os"
	"strings"
	"todo/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(ar *repositories.AuthRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtSecret := []byte(os.Getenv("JWT_SECRET"))

		//Получаем Authorization заголовок из запроса
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}
		// Разделяем содержимое заголовка и проверяем формат
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}
		//Сохраняем токен
		tokenStr := parts[1]

		//Декодируем и проверяем токен
		token, err := jwt.Parse(tokenStr,
			func(token *jwt.Token) (interface{}, error) {
				//Получаем инфу о методе шифрования токена
				_, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				//Возвращаем секрет для проверки подписи
				return jwtSecret, nil
			})

		//Проверяем результаты разкодирования токена и его валидность
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		//Получаем данные из токена и приводим их к читаемому виду
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		//Используем float64 потому что claims['id']возвращает тип interface{} который содержит float64
		//Преобразуем id пользователя из токена в int64 и ищем пользователя в базе данных
		uidFloat, ok := claims["id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}
		id := int64(uidFloat)
		user, err := ar.GetByID(id)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
