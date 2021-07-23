package usecase

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/ardafirdausr/posjoo-server/internal"
	"github.com/ardafirdausr/posjoo-server/internal/entity"
)

type AuthUsecase struct {
	userRepo     internal.UserRepository
	merchantRepo internal.MerchantRepository
	unitOfWork   internal.UnitOfWork
}

func NewAuthUsecase(userRepo internal.UserRepository, merchantRepo internal.MerchantRepository, unitOfWork internal.UnitOfWork) *AuthUsecase {
	uscase := new(AuthUsecase)
	uscase.userRepo = userRepo
	uscase.merchantRepo = merchantRepo
	uscase.unitOfWork = unitOfWork
	return uscase
}

func (uc AuthUsecase) Register(ctx context.Context, param entity.RegisterParam) (*entity.User, error) {
	existUser, err := uc.userRepo.GetUserByEmail(ctx, param.Email)
	_, errNotFound := err.(entity.ErrNotFound)
	if err != nil && !errNotFound {
		return nil, err
	}

	if existUser != nil && existUser.Email == param.Email {
		err := entity.ErrInvalidData{
			Message: "Email is already registered",
			Err:     errors.New("email is already registered"),
		}
		return nil, err
	}

	txCtx := uc.unitOfWork.Begin(ctx)
	createMerchantParam := entity.CreateMerchantParam{
		Name:      param.BusinessName,
		Address:   param.BusinessAddress,
		Phone:     param.BusinessPhone,
		CreatedAt: time.Now(),
	}
	merchant, err := uc.merchantRepo.CreateMerchant(txCtx, createMerchantParam)
	if err != nil {
		uc.unitOfWork.Rollback(txCtx)
		log.Println(err.Error())
		return nil, err
	}

	createUserPram := entity.CreateUserParam{
		Name:       param.Name,
		Email:      param.Email,
		Role:       entity.UserRoleOwner,
		Position:   "Owner",
		Password:   hashString(param.Password),
		MerchantID: merchant.ID,
		CreatedAt:  time.Now(),
	}
	user, err := uc.userRepo.CreateUser(txCtx, createUserPram)
	if err != nil {
		uc.unitOfWork.Rollback(txCtx)
		log.Println(err.Error())
		return nil, err
	}

	if err := uc.unitOfWork.Commit(txCtx); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return user, nil
}

func (uc AuthUsecase) GetUserFromCredential(ctx context.Context, param entity.LoginParam) (*entity.User, error) {
	errInvalid := entity.ErrInvalidData{
		Message: "Invalid Email or Password",
		Err:     errors.New("invalid email or password"),
	}

	user, err := uc.userRepo.GetUserByEmail(ctx, param.Email)
	if err != nil {
		log.Println(err.Error())
		return nil, errInvalid
	}

	hashPass := hashString(param.Password)
	if user.Password != hashPass {
		return nil, errInvalid
	}

	return user, nil
}

func (uc AuthUsecase) GetUserFromToken(ctx context.Context, token string, tokenizer internal.Tokenizer) (*entity.User, error) {
	if len(token) < 1 {
		return nil, errors.New("token is not provided")
	}

	payload, err := tokenizer.Parse(token)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("invalid token")
	}

	user, err := uc.userRepo.GetUserByID(ctx, payload.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (uc AuthUsecase) GenerateAuthToken(ctx context.Context, user entity.User, tokenizer internal.Tokenizer) (string, error) {
	tokenPayload := entity.TokenPayload{}
	tokenPayload.ID = user.ID
	tokenPayload.Name = user.Name
	tokenPayload.Email = user.Email
	tokenPayload.Role = user.Role
	tokenPayload.Position = user.Position
	tokenPayload.MerchantID = user.MerchantID
	tokenPayload.PhotoUrl = user.PhotoUrl
	token, err := tokenizer.Generate(tokenPayload)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return token, nil
}
