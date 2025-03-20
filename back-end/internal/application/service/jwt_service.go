package service

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"go.uber.org/zap"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"back-end/internal/infrastructure/logger"

)

type JWTService struct {
	publicKey *rsa.PublicKey
}


func NewJWTService(envKey string) (*JWTService, error) {
	pubKey, err := loadRSAPublicKeyFromEnv(envKey)
	if err != nil {
		return nil, err
	}
	return &JWTService{publicKey: pubKey}, nil
}

func loadRSAPublicKeyFromEnv(envKey string) (*rsa.PublicKey, error) {
	if err := godotenv.Load("/app/.bin/.env"); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	keyString := os.Getenv(envKey)
	if keyString == "" {
		return nil, fmt.Errorf("public key not found in environment")
	}

	block, _ := pem.Decode([]byte(keyString))
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DER encoded public key: %w", err)
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("public key is not RSA")
	}

	return rsaPub, nil
}

func (s *JWTService) VerifyToken(tokenString string) (*jwt.Token, error) {
	logger.Debug("Verifying JWT token", zap.String("token", tokenString))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			logger.Error("Unexpected signing method", zap.String("alg", fmt.Sprintf("%v", token.Header["alg"])))
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		logger.Debug("Token signature method validated")

		return s.publicKey, nil
	})

	if err != nil {
		logger.Error("JWT parsing failed", zap.Error(err))
	} else {
		logger.Debug("JWT successfully parsed")
	}

	return token, err
}