package internal

import "ardafirdausr/posjoo-server/internal/entity"

type UserUsecase interface {
	GetMerchantUsers(merchantID int64) ([]*entity.User, error)
	GetUser(userID int64) (*entity.User, error)
	CreateUser(param entity.CreateUserParam) (*entity.User, error)
	UpdateUser(userID int64, param entity.UpdateUserParam) (*entity.User, error)
	DeleteUser(userID int64) error
}

type ProductUsecase interface {
	GetMerchantProducts(merchantID int64) ([]*entity.Product, error)
	GetProduct(productID int64) (*entity.Product, error)
	CreateProduct(param entity.CreateProductParam) (*entity.Product, error)
	UpdateProduct(productID int64, param entity.UpdatedProductparam) (*entity.Product, error)
	UpdateProductPhoto(productID int64, param entity.UpdateProductPhotoParam) (*entity.Product, error)
	DeleteProduct(productID int64) error
}

type MerchantUsecase interface {
	CreateMerchant(param entity.CreateMerchantParam) (*entity.Merchant, error)
}
