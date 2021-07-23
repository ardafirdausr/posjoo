package app

import (
	"github.com/ardafirdausr/posjoo-server/internal"
	mysqlRepo "github.com/ardafirdausr/posjoo-server/internal/repository/mysql"
)

type repositories struct {
	userRepo     internal.UserRepository
	merchantRepo internal.MerchantRepository
	productRepo  internal.ProductRepository
}

func newRepositories(drivers *drivers) *repositories {
	userRepo := mysqlRepo.NewUserRepository(drivers.MySQL)
	merchantRepo := mysqlRepo.NewMerchantRepository(drivers.MySQL)
	productRepo := mysqlRepo.NewProductRepository(drivers.MySQL)

	return &repositories{
		userRepo:     userRepo,
		merchantRepo: merchantRepo,
		productRepo:  productRepo,
	}
}
