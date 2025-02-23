package utils

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenClaims struct {
	ID       string `json:"sub"`
	jwt.StandardClaims
}

func GenerateToken(userID string,ttl time.Duration, secretJWTKey string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	now := time.Now().UTC()

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims format")
	}
	claims["sub"] = userID
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	claims["jti"] = fmt.Sprintf("%d-%x", now.UnixNano(), generateRandomBytes(16))

	tokenString, err := token.SignedString([]byte(secretJWTKey))
	if err != nil {
		return "", fmt.Errorf("generating JWT Token failed: %w", err)
	}

	return tokenString, nil
}
func ValidateToken(tokenString string, secretKey string) (*TokenClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(secretKey), nil
    })

    if err != nil {
        return nil, fmt.Errorf("invalid token: %w", err)
    }

    if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, fmt.Errorf("invalid token claims")
}

func GenerateSecureOTP(length int) string {
	const charset = "0123456789"
	otp := make([]byte, length)
	randomBytes := generateRandomBytes(length)

	for i := 0; i < length; i++ {
		otp[i] = charset[int(randomBytes[i])%len(charset)]
	}

	return string(otp)
}

func generateRandomBytes(size int) []byte {
	randomBytes := make([]byte, size)
	_, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Printf("error generating random bytes: %v", err)
		return nil
	}
	return randomBytes
}
