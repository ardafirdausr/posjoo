package usecase

import (
	"errors"
	"log"
	"time"

	"github.com/ardafirdausr/posjoo-server/internal"
	"github.com/ardafirdausr/posjoo-server/internal/entity"
)

type UserUsecase struct {
	userRepo internal.UserRepository
}

func NewUserUsecase(userRepo internal.UserRepository) *UserUsecase {
	usecase := new(UserUsecase)
	usecase.userRepo = userRepo
	return usecase
}

func (uc *UserUsecase) GetMerchantUsers(merchantID int64) ([]*entity.User, error) {
	users, err := uc.userRepo.GetUsersByMerchantID(merchantID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return users, nil
}

func (uc *UserUsecase) GetUser(userID int64) (*entity.User, error) {
	user, err := uc.userRepo.GetUserByID(userID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return user, nil
}

func (uc *UserUsecase) CreateUser(param entity.CreateUserParam) (*entity.User, error) {
	existUser, err := uc.userRepo.GetUserByEmail(param.Email)
	_, errNotFound := err.(entity.ErrNotFound)
	if err != nil && !errNotFound {
		return nil, err
	}

	if existUser.Email == param.Email {
		err := entity.ErrInvalidData{
			Message: "Email is already registered",
			Err:     errors.New("email is already registered"),
		}
		return nil, err
	}

	param.Password = hashString(param.Password)
	param.CreatedAt = time.Now()
	user, err := uc.userRepo.CreateUser(param)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUsecase) UpdateUser(userID int64, param entity.UpdateUserParam) (*entity.User, error) {
	user, err := uc.userRepo.GetUserByID(userID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	existUser, err := uc.userRepo.GetUserByEmail(param.Email)
	_, errNotFound := err.(entity.ErrNotFound)
	if err != nil && !errNotFound {
		return nil, err
	}

	if existUser != nil && existUser.ID != user.ID && existUser.Email == param.Email {
		err := entity.ErrInvalidData{
			Message: "Email is already registered",
			Err:     errors.New("email is already registered"),
		}
		return nil, err
	}

	param.Password = hashString(param.Password)
	param.UpdatedAt = time.Now()
	if err = uc.userRepo.UpdateByID(userID, param); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return user, nil
}

func (uc *UserUsecase) DeleteUser(userID int64) error {
	if err := uc.userRepo.DeleteUserByID(userID); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
