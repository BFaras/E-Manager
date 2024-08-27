package middleware

import (
	"back-end/internal/application/service"
	"back-end/internal/infrastructure/logger"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")


        JWTService, err := service.NewJWTService("PUBLIC_KEY")
        if err != nil  {
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}

		token, err := JWTService.VerifyToken(tokenString)
		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		logger.Debug("Token parsed", zap.Reflect("Token", token))

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
		}

		userID, ok := claims["sub"].(string)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "User ID not found in token")
		}

		c.Set("userID", userID)
		logger.Debug("User ID extracted", zap.String("userId", userID))

		return next(c)
	}
}