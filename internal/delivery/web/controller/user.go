package controller

import (
	"log"
	"net/http"
	"strconv"

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

func (ctrl UserController) GetMerchantUsers(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*entity.JWTPayload)
	ctx := c.Request().Context()
	users, err := ctrl.ucs.UserUsecase.GetMerchantUsers(ctx, claims.MerchantID)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return jsonResponse(c, http.StatusOK, "Success", users)
}

func (ctrl UserController) GetUser(c echo.Context) error {
	userID, _ := strconv.ParseInt(c.Param("userID"), 10, 64)
	ctx := c.Request().Context()
	user, err := ctrl.ucs.UserUsecase.GetUser(ctx, userID)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return jsonResponse(c, http.StatusOK, "Success", user)
}

func (ctrl UserController) CreateUser(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*entity.JWTPayload)

	var param entity.CreateUserParam
	if err := c.Bind(&param); err != nil {
		log.Println(err.Error())
		return err
	}

	param.MerchantID = claims.MerchantID
	if err := c.Validate(&param); err != nil {
		log.Println(err.Error())
		return err
	}

	ctx := c.Request().Context()
	user, err := ctrl.ucs.UserUsecase.CreateUser(ctx, param)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return jsonResponse(c, http.StatusCreated, "Success", user)
}

func (ctrl UserController) UpdateUser(c echo.Context) error {
	userID, _ := strconv.ParseInt(c.Param("userID"), 10, 64)
	var param entity.UpdateUserParam
	if err := c.Bind(&param); err != nil {
		log.Println(err.Error())
		return err
	}

	if err := c.Validate(&param); err != nil {
		log.Println(err.Error())
		return err
	}

	ctx := c.Request().Context()
	user, err := ctrl.ucs.UserUsecase.UpdateUser(ctx, userID, param)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return jsonResponse(c, http.StatusOK, "Success", user)
}

func (ctrl UserController) UpdateUserPassword(c echo.Context) error {
	userID, _ := strconv.ParseInt(c.Param("userID"), 10, 64)
	var param entity.UpdateUserPasswordParam
	if err := c.Bind(&param); err != nil {
		log.Println(err.Error())
		return err
	}

	if err := c.Validate(&param); err != nil {
		log.Println(err.Error())
		return err
	}

	ctx := c.Request().Context()
	err := ctrl.ucs.UserUsecase.UpdateUserPassword(ctx, userID, param)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return jsonResponse(c, http.StatusOK, "Success", nil)
}

func (ctrl UserController) UpdateUserPhoto(c echo.Context) error {
	userID, _ := strconv.ParseInt(c.Param("userID"), 10, 64)
	fh, err := c.FormFile("photo")
	if err != nil {
		log.Println(err.Error())
		return echo.ErrBadRequest
	}

	ctx := c.Request().Context()
	user, err := ctrl.ucs.UserUsecase.UpdateUserPhoto(ctx, userID, fh)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return jsonResponse(c, http.StatusOK, "Success", user)
}

func (ctrl UserController) DeleteUser(c echo.Context) error {
	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		log.Println(err)
		eid := entity.ErrInvalidData{
			Message: "invalid user id",
			Err:     err,
		}
		return eid
	}

	ctx := c.Request().Context()
	if err := ctrl.ucs.UserUsecase.DeleteUser(ctx, userID); err != nil {
		log.Println(err.Error())
		return err
	}

	return jsonResponse(c, http.StatusNoContent, "Success", nil)
}
