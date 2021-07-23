package mysql

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/ardafirdausr/posjoo-server/internal/entity"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) *UserRepository {
	repo := new(UserRepository)
	repo.DB = DB
	return repo
}

func (repo *UserRepository) GetUserByID(ctx context.Context, userID int64) (*entity.User, error) {
	var query = "SELECT * FROM users WHERE id = ?"
	var row *sql.Row
	txKey := transactionContextKey("tx")
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		row = tx.QueryRow(query, userID)
	} else {
		row = repo.DB.QueryRowContext(ctx, query, userID)
	}

	var user entity.User
	var err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PhotoUrl,
		&user.Role,
		&user.Position,
		&user.Password,
		&user.MerchantID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		errNotFound := entity.ErrNotFound{
			Message: "User not found",
			Err:     err,
		}
		return nil, errNotFound
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var query = "SELECT * FROM users WHERE email = ?"
	var row *sql.Row
	txKey := transactionContextKey("tx")
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		row = tx.QueryRow(query, email)
	} else {
		row = repo.DB.QueryRowContext(ctx, query, email)
	}

	var user entity.User
	var err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PhotoUrl,
		&user.Role,
		&user.Position,
		&user.Password,
		&user.MerchantID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		errNotFound := entity.ErrNotFound{
			Message: "User not found",
			Err:     err,
		}
		return nil, errNotFound
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) GetUsersByMerchantID(ctx context.Context, merchantID int64) ([]*entity.User, error) {
	var query = "SELECT * FROM users WHERE merchant_id = ?"
	var rows *sql.Rows
	var err error
	txKey := transactionContextKey("tx")
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		rows, err = tx.Query(query, merchantID)
	} else {
		rows, err = repo.DB.QueryContext(ctx, query, merchantID)
	}

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	rows.Close()

	users := []*entity.User{}
	for rows.Next() {
		var user entity.User
		var err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.PhotoUrl,
			&user.Role,
			&user.Position,
			&user.Password,
			&user.MerchantID,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		users = append(users, &user)
	}
	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return users, nil
}

func (repo *UserRepository) CreateUser(ctx context.Context, param entity.CreateUserParam) (*entity.User, error) {
	var query = "INSERT INTO users(name, email, role, position, password, merchant_id, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"
	var res sql.Result
	var err error
	txKey := transactionContextKey("tx")
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		res, err = tx.Exec(query, param.Name, param.Email, param.Role, param.Position, param.Password, param.MerchantID, param.CreatedAt, param.CreatedAt)
	} else {
		res, err = repo.DB.ExecContext(ctx, query, param.Name, param.Email, param.Role, param.Position, param.Password, param.MerchantID, param.CreatedAt, param.CreatedAt)
	}

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	ID, err := res.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	user := &entity.User{
		ID:         ID,
		Name:       param.Name,
		Email:      param.Email,
		Role:       param.Role,
		Position:   param.Position,
		Password:   param.Password,
		MerchantID: param.MerchantID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return user, nil
}

func (repo *UserRepository) UpdateByID(ctx context.Context, userID int64, param entity.UpdateUserParam) error {
	var query = "UPDATE users SET name = ?, email = ?, role = ?, position = ?, password = ?, updated_at = ? WHERE id = ?"
	var res sql.Result
	var err error
	txKey := transactionContextKey("tx")
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		res, err = tx.Exec(query, param.Name, param.Email, param.Role, param.Position, param.Password, param.UpdatedAt, userID)
	} else {
		res, err = repo.DB.ExecContext(ctx, query, param.Name, param.Email, param.Role, param.Position, param.Password, param.UpdatedAt, userID)
	}

	if err != nil {
		log.Println(err.Error())
		return err
	}

	if num, _ := res.RowsAffected(); num < 1 {
		err := errors.New("failed to update user")
		enf := entity.ErrNotFound{
			Message: "User not found",
			Err:     err,
		}
		log.Println(enf.Error())
		return enf
	}

	return nil
}

func (repo *UserRepository) DeleteUserByID(ctx context.Context, userID int64) error {
	var query = "DELETE FROM users WHERE id = ?"
	var res sql.Result
	var err error
	txKey := transactionContextKey("tx")
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		res, err = tx.Exec(query, userID)
	} else {
		res, err = repo.DB.ExecContext(ctx, query, userID)
	}

	if err != nil {
		log.Println(err.Error())
		return err
	}

	if num, _ := res.RowsAffected(); num < 1 {
		err := errors.New("failed to delete user")
		enf := entity.ErrNotFound{
			Message: "User not found",
			Err:     err,
		}
		log.Println(enf.Error())
		return enf
	}

	return nil
}
