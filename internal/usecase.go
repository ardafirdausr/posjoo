package internal

import (
	"context"
	"mime/multipart"

	"github.com/ardafirdausr/posjoo-server/internal/entity"
)

type AuthUsecase interface {
	Register(ctx context.Context, param entity.RegisterParam) (*entity.User, error)
	GetUserFromToken(ctx context.Context, token string, tokenizer Tokenizer) (*entity.User, error)
	GetUserFromCredential(ctx context.Context, param entity.LoginParam) (*entity.User, error)
	GenerateAuthToken(ctx context.Context, user entity.User, tokenizer Tokenizer) (string, error)
}

type UserUsecase interface {
	GetMerchantUsers(ctx context.Context, merchantID int64) ([]*entity.User, error)
	GetUser(ctx context.Context, userID int64) (*entity.User, error)
	CreateUser(ctx context.Context, param entity.CreateUserParam) (*entity.User, error)
	UpdateUser(ctx context.Context, userID int64, param entity.UpdateUserParam) (*entity.User, error)
	UpdateUserPhoto(ctx context.Context, userID int64, photo *multipart.FileHeader) (*entity.User, error)
	UpdateUserPassword(ctx context.Context, userID int64, param entity.UpdateUserPasswordParam) error
	DeleteUser(ctx context.Context, userID int64) error
}

type ProductUsecase interface {
	GetMerchantProducts(ctx context.Context, merchantID int64) ([]*entity.Product, error)
	GetProduct(ctx context.Context, productID int64) (*entity.Product, error)
	CreateProduct(ctx context.Context, param entity.CreateProductParam) (*entity.Product, error)
	UpdateProduct(ctx context.Context, productID int64, param entity.UpdatedProductparam) (*entity.Product, error)
	UpdateProductPhoto(ctx context.Context, productID int64, photo *multipart.FileHeader) (*entity.Product, error)
	DeleteProduct(ctx context.Context, productID int64) error
}

type MerchantUsecase interface {
	CreateMerchant(ctx context.Context, param entity.CreateMerchantParam) (*entity.Merchant, error)
}
