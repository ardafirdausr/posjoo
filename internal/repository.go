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
	UpdateByID(ctx context.Context, userID int64, param entity.UpdateUserParam) error
	DeleteUserByID(ctx context.Context, userID int64) error
}

type MerchantRepository interface {
	CreateMerchant(ctx context.Context, param entity.CreateMerchantParam) (*entity.Merchant, error)
}

type ProductRepository interface {
	GetProductByID(ctx context.Context, productID int64) (*entity.Product, error)
	GetProductBySKU(ctx context.Context, SKU string) (*entity.Product, error)
	GetProductsByMerchantID(ctx context.Context, merchantID int64) ([]*entity.Product, error)
	CreateProduct(ctx context.Context, param entity.CreateProductParam) (*entity.Product, error)
	UpdateProductByID(ctx context.Context, productId int64, param entity.UpdatedProductparam) error
	DeleteProductByID(ctx context.Context, productId int64) error
}
