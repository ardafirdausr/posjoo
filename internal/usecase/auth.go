package usecase

import (
	"errors"
	"log"
	"time"

	"github.com/ardafirdausr/posjoo-server/internal"
	"github.com/ardafirdausr/posjoo-server/internal/entity"
)

type AuthUsecase struct {
	userRepo internal.UserRepository
}

func NewAuthUsecase(userRepo internal.UserRepository) *AuthUsecase {
	return &AuthUsecase{userRepo: userRepo}
}

func (uc AuthUsecase) Register(param entity.RegisterParam) (*entity.User, error) {
	existUser, err := uc.userRepo.GetUserByEmail(param.Email)
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

	createUserPram := entity.CreateUserParam{
		Name:       param.Name,
		Email:      param.Email,
		Role:       entity.UserRoleOwner,
		Position:   "Owner",
		Password:   hashString(param.Password),
		MerchantID: 0,
		CreatedAt:  time.Now(),
	}
	user, err := uc.userRepo.CreateUser(createUserPram)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc AuthUsecase) GetUserFromCredential(param entity.LoginParam) (*entity.User, error) {
	errInvalid := entity.ErrInvalidData{
		Message: "Invalid Email or Password",
		Err:     errors.New("invalid email or password"),
	}

	user, err := uc.userRepo.GetUserByEmail(param.Email)
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

func (uc AuthUsecase) GetUserFromToken(token string, tokenizer internal.Tokenizer) (*entity.User, error) {
	if len(token) < 1 {
		return nil, errors.New("token is not provided")
	}

	payload, err := tokenizer.Parse(token)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("invalid token")
	}

	user, err := uc.userRepo.GetUserByID(payload.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (uc AuthUsecase) GenerateAuthToken(user entity.User, tokenizer internal.Tokenizer) (string, error) {
	tokenPayload := entity.TokenPayload{}
	tokenPayload.ID = user.ID
	tokenPayload.Name = user.Name
	tokenPayload.Email = user.Email
	tokenPayload.PhotoUrl = user.PhotoUrl
	token, err := tokenizer.Generate(tokenPayload)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return token, nil
}
