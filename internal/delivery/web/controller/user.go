package controller

import (
	"github.com/ardafirdausr/posjoo-server/internal/app"
	"github.com/ardafirdausr/posjoo-server/internal/entity"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	ucs *app.Usecases
}

func NewUserController(ucs *app.Usecases) *UserController {
	userController := new(UserController)
	userController.ucs = ucs
	return userController
}

func (ctrl UserController) GetUsers(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*entity.JWTPayload)
	ctx := c.Request().Context()
	ctrl.ucs.UserUsecase.GetMerchantUsers(ctx, claims.ExpiresAt)

	return nil
}

func (ctrl UserController) CreateUser(c echo.Context) error {
	return nil
}

func (ctrl UserController) UpdateUser(c echo.Context) error {
	return nil
}

func (ctrl UserController) DeleteUser(c echo.Context) error {
	return nil
}
