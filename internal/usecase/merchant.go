package usecase

import (
	"log"

	"github.com/ardafirdausr/posjoo-server/internal"
	"github.com/ardafirdausr/posjoo-server/internal/entity"
)

type MerchantUsecase struct {
	merchantRepo internal.MerchantRepository
}

func NewMerchantUsecase(merchantRepo internal.MerchantRepository) *MerchantUsecase {
	usecase := new(MerchantUsecase)
	usecase.merchantRepo = merchantRepo
	return usecase
}

func (uc MerchantUsecase) CreateMerchant(param entity.CreateMerchantParam) (*entity.Merchant, error) {
	merchant, err := uc.merchantRepo.CreateMerchant(param)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return merchant, nil
}
