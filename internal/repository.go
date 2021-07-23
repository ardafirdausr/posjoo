package internal

import "github.com/ardafirdausr/posjoo-server/internal/entity"

type UserRepository interface {
	GetUserByID(userID int64) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetUsersByMerchantID(merchantID int64) ([]*entity.User, error)
	CreateUser(param entity.CreateUserParam) (*entity.User, error)
	UpdateByID(userID int64, param entity.UpdateUserParam) error
	DeleteUserByID(userID int64) error
}

type MerchantRepository interface {
	CreateMerchant(param entity.CreateMerchantParam) (*entity.Merchant, error)
}

type ProductRepository interface {
	GetProductByID(productID int64) (*entity.Product, error)
	GetProductBySKU(SKU string) (*entity.Product, error)
	GetProductsByMerchantID(merchantID int64) ([]*entity.Product, error)
	CreateProduct(param entity.CreateProductParam) (*entity.Product, error)
	UpdateProductByID(productId int64, param entity.UpdatedProductparam) error
	DeleteProductByID(productId int64) error
}
