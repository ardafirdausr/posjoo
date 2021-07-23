package usecase

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/ardafirdausr/posjoo-server/internal"
	"github.com/ardafirdausr/posjoo-server/internal/entity"
)

type ProductUsecase struct {
	ProductRepo internal.ProductRepository
}

func NewProductUsecase(ProductRepo internal.ProductRepository) *ProductUsecase {
	usecase := new(ProductUsecase)
	usecase.ProductRepo = ProductRepo
	return usecase
}

func (uc ProductUsecase) GetMerchantProducts(ctx context.Context, merchantID int64) ([]*entity.Product, error) {
	products, err := uc.ProductRepo.GetProductsByMerchantID(ctx, merchantID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return products, err
}

func (uc ProductUsecase) GetProduct(ctx context.Context, productID int64) (*entity.Product, error) {
	product, err := uc.ProductRepo.GetProductByID(ctx, productID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return product, err
}

func (uc ProductUsecase) CreateProduct(ctx context.Context, param entity.CreateProductParam) (*entity.Product, error) {
	existProduct, err := uc.ProductRepo.GetProductBySKU(ctx, param.SKU)
	_, errNotFound := err.(entity.ErrNotFound)
	if err != nil && !errNotFound {
		log.Println(err.Error())
		return nil, err
	}

	if existProduct != nil && existProduct.SKU == param.SKU {
		err := entity.ErrInvalidData{
			Message: "SKU is already registered",
			Err:     errors.New("SKU is already registered"),
		}
		log.Println(err.Error())
		return nil, err
	}

	param.CreatedAt = time.Now()
	product, err := uc.ProductRepo.CreateProduct(ctx, param)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return product, nil
}

func (uc ProductUsecase) UpdateProduct(ctx context.Context, productID int64, param entity.UpdatedProductparam) (*entity.Product, error) {
	product, err := uc.ProductRepo.GetProductByID(ctx, productID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	existProduct, err := uc.ProductRepo.GetProductBySKU(ctx, param.SKU)
	_, errNotFound := err.(entity.ErrNotFound)
	if err != nil && !errNotFound {
		log.Println(err.Error())
		return nil, err
	}

	if existProduct != nil && existProduct.ID != product.ID && existProduct.SKU == param.SKU {
		err := entity.ErrInvalidData{
			Message: "SKU is already registered",
			Err:     errors.New("SKU is already registered"),
		}
		log.Println(err.Error())
		return nil, err
	}

	param.UpdatedAt = time.Now()
	err = uc.ProductRepo.UpdateProductByID(ctx, productID, param)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return uc.ProductRepo.GetProductByID(ctx, productID)
}

func (uc ProductUsecase) UpdateProductPhoto(ctx context.Context, productID int64, param entity.UpdateProductPhotoParam) (*entity.Product, error) {
	return nil, nil
}

func (uc ProductUsecase) DeleteProduct(ctx context.Context, productID int64) error {
	if err := uc.ProductRepo.DeleteProductByID(ctx, productID); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
