package app

import (
	"github.com/ardafirdausr/posjoo-server/internal"
	"github.com/ardafirdausr/posjoo-server/internal/usecase"
)

type Usecases struct {
	AuthUsecase    internal.AuthUsecase
	UserUsecase    internal.UserUsecase
	ProductUsecase internal.ProductUsecase
}

func newUsecases(repos *repositories, services *services) *Usecases {
	usecases := new(Usecases)
	usecases.AuthUsecase = usecase.NewAuthUsecase(repos.userRepo, repos.merchantRepo, repos.unitOfWork)
	usecases.UserUsecase = usecase.NewUserUsecase(repos.userRepo, services.storageService)
	usecases.ProductUsecase = usecase.NewProductUsecase(repos.productRepo, services.storageService)
	return usecases
}
