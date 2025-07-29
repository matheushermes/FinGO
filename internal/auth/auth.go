package auth

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/matheushermes/FinGO/configs"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrNoToken      = errors.New("authorization token not found")
)

func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{
		"authorized": true,
		"exp":        time.Now().Add(8 * time.Hour).Unix(),
		"userId":     userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(configs.SECRET_KEY))
}

func ExtractToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}
	return ""
}

func returnVericationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return configs.SECRET_KEY, nil
}

func ValidateToken(c *gin.Context) error {
	tokenString := ExtractToken(c)
	if tokenString == "" {
		return ErrNoToken
	}

	token, err := jwt.Parse(tokenString, returnVericationKey)
	if err != nil {
		return fmt.Errorf("token parsing failed: %w", err)
	}

	if !token.Valid {
		return ErrInvalidToken
	}

	return nil
}

func ExtractDataFromToken(c *gin.Context) (uint64, error) {
	tokenString := ExtractToken(c)
	if tokenString == "" {
		return 0, ErrNoToken
	}

	token, err := jwt.Parse(tokenString, returnVericationKey)
	if err != nil || !token.Valid {
		return 0, ErrInvalidToken
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok {
		userIDStr := fmt.Sprintf("%.0f", permissions["userId"])
		userID, err := strconv.ParseUint(userIDStr, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("failed to parse userID: %w", err)
		}
		return userID, nil
	}

	return 0, ErrInvalidToken
}