package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"fmt"
    "net/http"
    "strings"
	"github.com/dgrijalva/jwt-go"
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

var secretKey = []byte("your-secret-key") 

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        tokenString := c.Request().Header.Get("Authorization")
        if tokenString == "" {
            return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
        }

        tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return secretKey, nil
        })

        if err != nil || !token.Valid {
            return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
        }

        userID, ok := claims["sub"].(string)
        if !ok {
            return echo.NewHTTPError(http.StatusUnauthorized, "User ID not found in token")
        }

        c.Set("userID", userID)

        return next(c)
    }
}

func ProtectedHandler(c echo.Context) error {
    userID := c.Get("userID").(string)
    return c.String(http.StatusOK, fmt.Sprintf("Hello, user %s!", userID))
}
