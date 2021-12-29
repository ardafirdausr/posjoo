package mysql

import (
	"context"
	"database/sql"
	"errors"
)

type transactionContextKey string

type MySQLUnitOfWork struct {
	DB *sql.DB
}

func NewMySQLUnitOfWork(DB *sql.DB) *MySQLUnitOfWork {
	unitOfWork := new(MySQLUnitOfWork)
	unitOfWork.DB = DB
	return unitOfWork
}

func (uw MySQLUnitOfWork) Begin(ctx context.Context) context.Context {
	transactionKey := transactionContextKey("tx")
	tx, _ := uw.DB.BeginTx(ctx, nil)
	transactionContext := context.WithValue(ctx, transactionKey, tx)
	return transactionContext
}

func (uw MySQLUnitOfWork) Commit(ctx context.Context) error {
	transactionKey := transactionContextKey("tx")
	tx, ok := ctx.Value(transactionKey).(*sql.Tx)
	if !ok {
		return errors.New("failed get transaction")
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (uw MySQLUnitOfWork) Rollback(ctx context.Context) error {
	tx, ok := ctx.Value("tx").(*sql.Tx)
	if !ok {
		return errors.New("failed get transaction")
	}

	if err := tx.Rollback(); err != nil {
		return err
	}

	return nil
}
