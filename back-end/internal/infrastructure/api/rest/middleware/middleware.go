package middleware

import (
	"net/http"
	"strings"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
	"github.com/MicahParks/keyfunc"
	"go.uber.org/zap"
)

type Middleware struct {
}

func New() *Middleware {
	return &Middleware{}
}

func (m *Middleware) CORSConfig() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:3000",
            "http://localhost:3001",
			"*",
		},
		AllowMethods: []string{
			echo.GET,
            echo.PATCH,
			echo.POST,
			echo.PUT,
			echo.DELETE,
		},
		AllowHeaders: []string{
			echo.HeaderContentType,
			echo.HeaderAuthorization,
		},
		AllowCredentials: false,
	})
}

func (m *Middleware) JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	jwksURL := "https://fast-barnacle-55.clerk.accounts.dev/.well-known/jwks.json" 

	options := keyfunc.Options{
		RefreshInterval: time.Hour,
		RefreshErrorHandler: func(err error) {
			zap.L().Error("Error refreshing JWKS", zap.Error(err))
		},
		RefreshTimeout: 30 * time.Second,
	}

	jwks, err := keyfunc.Get(jwksURL, options)
	if err != nil {
		zap.L().Fatal("Failed to get JWKS from Clerk", zap.Error(err))
	}

	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		zap.L().Debug("Received Authorization Header", zap.String("Authorization", tokenString))
		if tokenString == "" {
			zap.L().Error("No token found in request headers")
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		zap.L().Debug("Extracted Token", zap.String("token", tokenString))

		token, err := jwt.Parse(tokenString, jwks.Keyfunc)
		if err != nil || !token.Valid {
			zap.L().Error("Token verification failed", zap.Error(err))
			if err != nil {
				zap.L().Debug("Token error details", zap.Any("error", err))
			}
			var ve *jwt.ValidationError
			if errors.As(err, &ve) && ve.Errors&jwt.ValidationErrorExpired != 0 {
				zap.L().Error("Token has expired", zap.Error(err))
				return echo.NewHTTPError(http.StatusUnauthorized, "Token expired")
			}
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			zap.L().Error("Failed to extract token claims")
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
		}


		zap.L().Debug("Token Claims", zap.Any("claims", claims))
		if exp, ok := claims["exp"].(float64); ok {
			currentTime := float64(time.Now().Unix())
			zap.L().Debug("Token expiration", zap.Float64("exp", exp), zap.Float64("currentTime", currentTime))
			zap.L().Debug("Time until expiration", zap.Float64("secondsLeft", exp-currentTime))
		}
		if iat, ok := claims["iat"].(float64); ok {
			zap.L().Debug("Token issued at", zap.Float64("iat", iat))
		}
		if nbf, ok := claims["nbf"].(float64); ok {
			zap.L().Debug("Token not valid before", zap.Float64("nbf", nbf))
		}


		userID, ok := claims["sub"].(string)
		if !ok {
			zap.L().Error("User ID not found in token", zap.Any("claims", claims))
			return echo.NewHTTPError(http.StatusUnauthorized, "User ID not found in token")
		}

		c.Set("userID", userID)
		zap.L().Debug("User ID successfully extracted", zap.String("userId", userID))

		return next(c)
	}
}