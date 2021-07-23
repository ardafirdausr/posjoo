package controller

import (
	"log"
	"net/http"
	"os"

	"github.com/ardafirdausr/posjoo-server/internal/app"
	"github.com/ardafirdausr/posjoo-server/internal/entity"
	"github.com/ardafirdausr/posjoo-server/internal/service/token"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	ucs *app.Usecases
}

func NewAuthController(ucs *app.Usecases) *AuthController {
	return &AuthController{ucs: ucs}
}

func (ctrl AuthController) Register(c echo.Context) error {
	return nil
}

func (ctrl AuthController) Login(c echo.Context) error {
	var param entity.LoginParam
	if err := c.Bind(&param); err != nil {
		return echo.ErrInternalServerError
	}

	user, err := ctrl.ucs.AuthUsecase.GetUserFromCredential(param)
	if err != nil {
		log.Panicln(err.Error())
		return err
	}

	JWTSecretKey := os.Getenv("JWT_SECRET_KEY")
	JWTToknizer := token.NewJWTTokenizer(JWTSecretKey)
	JWTToken, err := ctrl.ucs.AuthUsecase.GenerateAuthToken(*user, JWTToknizer)
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
