package middleware

import (
	"github.com/ardafirdausr/posjoo-server/internal/entity"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func OwnerAccess() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			claims := token.Claims.(*entity.JWTPayload)
			if claims.Role == entity.UserRoleOwner {
				return next(c)
			}

			return echo.ErrForbidden
		}
	}
}
