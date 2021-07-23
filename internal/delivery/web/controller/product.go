package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ardafirdausr/posjoo-server/internal/app"
	"github.com/ardafirdausr/posjoo-server/internal/entity"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	ucs *app.Usecases
}

func NewProductController(ucs *app.Usecases) *ProductController {
	productController := new(ProductController)
	productController.ucs = ucs
	return productController
}

func (ctrl ProductController) GetMerchantProducts(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*entity.JWTPayload)
	ctx := c.Request().Context()
	products, err := ctrl.ucs.ProductUsecase.GetMerchantProducts(ctx, claims.MerchantID)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return jsonResponse(c, http.StatusOK, "Success", products)
}

func (ctrl ProductController) GetProduct(c echo.Context) error {
	productID, _ := strconv.ParseInt(c.Param("productID"), 10, 64)
	ctx := c.Request().Context()
	product, err := ctrl.ucs.ProductUsecase.GetProduct(ctx, productID)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return jsonResponse(c, http.StatusOK, "Success", product)
}

func (ctrl ProductController) CreateProduct(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*entity.JWTPayload)

	var param entity.CreateProductParam
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
	product, err := ctrl.ucs.ProductUsecase.CreateProduct(ctx, param)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return jsonResponse(c, http.StatusCreated, "Success", product)
}

func (ctrl ProductController) UpdateProduct(c echo.Context) error {
	productID, _ := strconv.ParseInt(c.Param("productID"), 10, 64)
	var param entity.UpdatedProductparam
	if err := c.Bind(&param); err != nil {
		log.Println(err.Error())
		return err
	}

	if err := c.Validate(&param); err != nil {
		log.Println(err.Error())
		return err
	}

	ctx := c.Request().Context()
	product, err := ctrl.ucs.ProductUsecase.UpdateProduct(ctx, productID, param)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return jsonResponse(c, http.StatusOK, "Success", product)
}

func (ctrl ProductController) UpdateProductPhoto(c echo.Context) error {
	productID, _ := strconv.ParseInt(c.Param("productID"), 10, 64)
	var param entity.UpdatedProductparam
	if err := c.Bind(&param); err != nil {
		log.Println(err.Error())
		return err
	}

	fh, err := c.FormFile("photo")
	if err != nil {
		log.Println(err.Error())
		return echo.ErrBadRequest
	}

	if fh == nil {
		return echo.ErrBadRequest
	}

	rule := map[string]int64{
		".jpg":  1024 * 1000 * 4,
		".jpeg": 1024 * 1000 * 4,
		".png":  1024 * 1000 * 4,
	}
	photoExt := strings.ToLower(filepath.Ext(fh.Filename))
	maxSize, ok := rule[photoExt]
	if !ok {
		return entity.ErrInvalidData{
			Message: "photo extension must be .jpg, .jpeg, or .png",
			Err:     errors.New("photo extension must be .jpg, .jpeg, or .png"),
		}
	}

	if fh.Size > maxSize {
		return entity.ErrInvalidData{
			Message: "Max photo size is 4MB",
			Err:     errors.New("max photo size is 4MB"),
		}
	}

	fmt.Println(fh.Size)

	ctx := c.Request().Context()
	user, err := ctrl.ucs.ProductUsecase.UpdateProductPhoto(ctx, productID, fh)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return jsonResponse(c, http.StatusOK, "Success", user)
}

func (ctrl ProductController) DeleteProduct(c echo.Context) error {
	productID, err := strconv.ParseInt(c.Param("productID"), 10, 64)
	if err != nil {
		log.Println(err)
		eid := entity.ErrInvalidData{
			Message: "invalid product id",
			Err:     err,
		}
		return eid
	}

	ctx := c.Request().Context()
	if err := ctrl.ucs.ProductUsecase.DeleteProduct(ctx, productID); err != nil {
		log.Println(err.Error())
		return err
	}

	return jsonResponse(c, http.StatusNoContent, "Success", nil)
}
