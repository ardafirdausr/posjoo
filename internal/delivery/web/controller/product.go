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
