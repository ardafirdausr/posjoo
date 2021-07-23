package mysql

import (
	"context"
	"database/sql"
	"log"

	"github.com/ardafirdausr/posjoo-server/internal/entity"
)

type MerchantRepository struct {
	DB *sql.DB
}

func NewMerchantRepository(DB *sql.DB) *MerchantRepository {
	repo := new(MerchantRepository)
	repo.DB = DB
	return repo
}

func (repo MerchantRepository) CreateMerchant(param entity.CreateMerchantParam) (*entity.Merchant, error) {
	ctx := context.TODO()
	res, err := repo.DB.ExecContext(
		ctx,
		"INSERT INTO merchants(name, address, phone, created_at, updated_at) VALUES(?, ?, ?, ?, ?)",
		param.Name, param.Address, param.Phone, param.CreatedAt, param.CreatedAt)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	merchant := &entity.Merchant{
		ID:        ID,
		Name:      param.Name,
		Address:   param.Address,
		Phone:     param.Phone,
		CreatedAt: param.CreatedAt,
		UpdatedAt: param.CreatedAt,
	}

	return merchant, nil
}
