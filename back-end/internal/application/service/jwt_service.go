package service

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
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
	if err := godotenv.Load("../../.bin/.env"); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	keyString := os.Getenv(envKey)
	if keyString == "" {
		return nil, fmt.Errorf("public key not found in environment")
	}

	logger.Debug("Loaded key string", zap.String("keyString", keyString))

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
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.publicKey, nil
	})
	return token, err
}