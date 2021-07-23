package app

import (
	"github.com/ardafirdausr/posjoo-server/internal"
	mysqlRepo "github.com/ardafirdausr/posjoo-server/internal/repository/mysql"
)

type repositories struct {
	userRepo     internal.UserRepository
	merchantRepo internal.MerchantRepository
	productRepo  internal.ProductRepository
	unitOfWork   internal.UnitOfWork
}

func newRepositories(drivers *drivers) *repositories {
	repos := new(repositories)
	repos.userRepo = mysqlRepo.NewUserRepository(drivers.MySQL)
	repos.merchantRepo = mysqlRepo.NewMerchantRepository(drivers.MySQL)
	repos.productRepo = mysqlRepo.NewProductRepository(drivers.MySQL)
	repos.unitOfWork = mysqlRepo.NewMySQLUnitOfWork(drivers.MySQL)
	return repos
}
