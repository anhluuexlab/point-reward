package middlewares

import (
	"github.com/zett-8/go-clean-echo/security"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zett-8/go-clean-echo/models"
)

func JwtMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &models.JwtCustomClaims{},
		SigningKey: []byte(security.JWT_KEY),
	}

	return middleware.JWTWithConfig(config)
}
