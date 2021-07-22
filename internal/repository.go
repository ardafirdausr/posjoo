package internal

import "ardafirdausr/posjoo-server/internal/entity"

type UserRepository interface {
	GetUserByID(ID int64) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetUsersByMerchantID(merchantID int64) ([]*entity.User, error)
	CreateUser(param entity.CreateUserParam) (*entity.Product, error)
	UpdateByID(ID int64, param entity.UpdateUserParam) error
	DeleteUserByID(ID int64) error
}

type MerchantRepository interface {
	CreateMerchant(param entity.CreateMerchantParam) (*entity.Merchant, error)
}

type ProductRepository interface {
	GetProductByID(productID int64) (*entity.Product, error)
	GetProductsByMerchantID(merchantID int64) ([]*entity.Product, error)
	CreteProduct(param entity.CreateProductParam) (*entity.Product, error)
	UpdateProductByID(productId int64, param entity.UpdatedProductparam) error
	DeleteProductByID(productId int64) error
}
