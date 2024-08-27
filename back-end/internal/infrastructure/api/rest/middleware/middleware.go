package middleware

import (
	"back-end/internal/infrastructure/logger"
	"fmt"
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

var publicKey = []byte("pk_test_ZmFzdC1iYXJuYWNsZS01NS5jbGVyay5hY2NvdW50cy5kZXYk") 

func (m *Middleware) JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        logger.Debug("Got to the JWT")

        // Extract token from header
        tokenString := c.Request().Header.Get("Authorization")
        if tokenString == "" {
            return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
        }

        // Remove "Bearer " prefix
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        // Parse the token
        token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // Ensure token method is RS256
            if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return publicKey, nil
        })

        /*@Todo@: Add token validation should be   err != nil || !token.Valid  and not err !*/

        logger.Debug("Token parsed", zap.Reflect("Token", token))
        /*
        if err != nil || !token.Valid {
            return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
        }/***/

        // Extract claims
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
        }

        // Extract user ID from claims
        userID, ok := claims["sub"].(string)
        if !ok {
            return echo.NewHTTPError(http.StatusUnauthorized, "User ID not found in token")
        }

        // Set user ID in context
        c.Set("userID", userID)
        logger.Debug("User ID extracted", zap.String("userId", userID))

        return next(c)
    }
}

func ProtectedHandler(c echo.Context) error {
    userID := c.Get("userID").(string)
    return c.String(http.StatusOK, fmt.Sprintf("Hello, user %s!", userID))
}
