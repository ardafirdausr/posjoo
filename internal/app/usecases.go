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

func newUsecases(repos *repositories) *Usecases {
	authUsecase := usecase.NewAuthUsecase(repos.userRepo)
	userUsecase := usecase.NewUserUsecase(repos.userRepo)
	productUsecase := usecase.NewProductUsecase(repos.productRepo)
	return &Usecases{
		AuthUsecase:    authUsecase,
		UserUsecase:    userUsecase,
		ProductUsecase: productUsecase,
	}
}
