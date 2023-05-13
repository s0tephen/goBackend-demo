package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(username string) (string, error) {
	// Get secret key from environment variable
	secret := viper.GetString("jwt.secret_key")
	if secret == "" {
		return "", errors.New("JWT_SECRET_KEY not set")
	}

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    viper.GetString("jwt.Issuer"),
		},
	}
	// Generate encoded token and send it as response.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ParseToken 解析token
func ParseToken(tokenString string, secret []byte) (*Claims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	// Extract the claims
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

// VerifyToken 验证token
func VerifyToken(tokenString string, secret []byte) (*Claims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return secret key for validation
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	// Check token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Check standard claims
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	if err := claims.Valid(); err != nil {
		return nil, err
	}

	// Check custom claims
	if claims.Username == "" {
		return nil, fmt.Errorf("missing username claim")
	}
	return claims, nil
}
