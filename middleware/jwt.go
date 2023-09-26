package middleware

import (
	"GitTrendy/model"
	"GitTrendy/security"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims: &model.JwtCustomClaims{},
		SigningKey: security.SECRET_KEY,
	}

	return middleware.JWTWithConfig(config)
}
