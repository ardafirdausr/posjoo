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

func (repo MerchantRepository) CreateMerchant(ctx context.Context, param entity.CreateMerchantParam) (*entity.Merchant, error) {
	var query = "INSERT INTO merchants(name, address, phone, created_at, updated_at) VALUES(?, ?, ?, ?, ?)"
	var res sql.Result
	var err error
	txKey := transactionContextKey("tx")
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		res, err = tx.Exec(query, param.Name, param.Address, param.Phone, param.CreatedAt, param.CreatedAt)
	} else {
		res, err = repo.DB.ExecContext(ctx, query, param.Name, param.Address, param.Phone, param.CreatedAt, param.CreatedAt)
	}

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
