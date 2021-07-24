package internal

import (
	"context"

	"github.com/ardafirdausr/posjoo-server/internal/entity"
)

type UnitOfWork interface {
	Begin(context.Context) context.Context
	Commit(context.Context) error
	Rollback(context.Context) error
}

type UserRepository interface {
	GetUserByID(ctx context.Context, userID int64) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUsersByMerchantID(ctx context.Context, merchantID int64) ([]*entity.User, error)
	CreateUser(ctx context.Context, param entity.CreateUserParam) (*entity.User, error)
	UpdateUserByID(ctx context.Context, userID int64, param entity.UpdateUserParam) error
	UpdateUserPasswordByID(ctx context.Context, userID int64, password string) error
	UpdateUserPhotoByID(ctx context.Context, userID int64, url string) error
	DeleteUserByID(ctx context.Context, userID int64) error
}

type MerchantRepository interface {
	CreateMerchant(ctx context.Context, param entity.CreateMerchantParam) (*entity.Merchant, error)
}

type ProductRepository interface {
	GetProductByID(ctx context.Context, productID int64) (*entity.Product, error)
	GetProductBySKUIndex(ctx context.Context, merchantID int64, SKU string) (*entity.Product, error)
	GetProductsByMerchantID(ctx context.Context, merchantID int64) ([]*entity.Product, error)
	CreateProduct(ctx context.Context, param entity.CreateProductParam) (*entity.Product, error)
	UpdateProductByID(ctx context.Context, productID int64, param entity.UpdatedProductparam) error
	UpdateProductPhotoByID(ctx context.Context, productID int64, url string) error
	DeleteProductByID(ctx context.Context, productID int64) error
}
