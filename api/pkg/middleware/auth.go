package middleware

import (
	"api/pkg/log"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"api/pkg/utils"
)

var secretKey []byte

func init() {
	if err := godotenv.Load(); err != nil {
		log.Error("Error loading .env file", err)
	}

	envKey := os.Getenv("JWT_SECRET_KEY")
	if envKey == "" {
		log.Error("JWT_SECRET_KEY environment variable is not set, using default secret key", nil)
	} else {
		secretKey = []byte(envKey)
	}
}

// APIAuthMiddleware is the middleware function for JWT authentication
func APIAuthMiddleware(c *fiber.Ctx) error {
	// Token'ı başlıkta ara
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		log.JWTError("Unauthorized access attempt: No token provided", nil, c)
		return utils.JSONError(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Token'ı doğrula
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// Hatalı imzalama yöntemi
			errMsg := fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])
			log.JWTError(errMsg, nil, c)
			return nil, fmt.Errorf(errMsg)
		}
		return secretKey, nil
	})

	// Eğer token geçerli değilse, hata döndür
	if err != nil || !token.Valid {
		log.JWTError("Unauthorized access attempt: Invalid token", err, c)
		return utils.JSONError(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	log.JWTInfo("Authorized request", c) // Başarılı yetkilendirme

	// Token geçerliyse, sonraki handler'a geç
	return c.Next()
}
