package controller

import (
	"log"
	"net/http"
	"os"

	"github.com/ardafirdausr/posjoo-server/internal/app"
	"github.com/ardafirdausr/posjoo-server/internal/entity"
	"github.com/ardafirdausr/posjoo-server/internal/pkg/token"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	ucs *app.Usecases
}

func NewAuthController(ucs *app.Usecases) *AuthController {
	return &AuthController{ucs: ucs}
}

func (ctrl AuthController) Register(c echo.Context) error {
	var param entity.RegisterParam
	if err := c.Bind(&param); err != nil {
		return echo.ErrInternalServerError
	}

	if err := c.Validate(&param); err != nil {
		log.Println(err.Error())
		return err
	}

	ctx := c.Request().Context()
	user, err := ctrl.ucs.AuthUsecase.Register(ctx, param)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	JWTSecretKey := os.Getenv("JWT_SECRET_KEY")
	JWTToknizer := token.NewJWTTokenizer(JWTSecretKey)
	JWTToken, err := ctrl.ucs.AuthUsecase.GenerateAuthToken(ctx, *user, JWTToknizer)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	response := echo.Map{
		"message": "Login Successful",
		"data":    user,
		"token":   JWTToken,
	}
	return c.JSON(http.StatusOK, response)
}

func (ctrl AuthController) Login(c echo.Context) error {
	var param entity.LoginParam
	if err := c.Bind(&param); err != nil {
		log.Println(err.Error())
		return echo.ErrInternalServerError
	}

	if err := c.Validate(&param); err != nil {
		log.Println(err.Error())
		return err
	}

	ctx := c.Request().Context()
	user, err := ctrl.ucs.AuthUsecase.GetUserFromCredential(ctx, param)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	JWTSecretKey := os.Getenv("JWT_SECRET_KEY")
	JWTToknizer := token.NewJWTTokenizer(JWTSecretKey)
	JWTToken, err := ctrl.ucs.AuthUsecase.GenerateAuthToken(ctx, *user, JWTToknizer)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	response := echo.Map{
		"message": "Login Successful",
		"data":    user,
		"token":   JWTToken,
	}
	return c.JSON(http.StatusOK, response)
}
