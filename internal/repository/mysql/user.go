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

func (repo *UserRepository) GetUserByID(userID int64) (*entity.User, error) {
	ctx := context.TODO()
	row := repo.DB.QueryRowContext(ctx, "SELECT * FROM users WHERE id = ?", userID)

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

func (repo *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	ctx := context.TODO()
	row := repo.DB.QueryRowContext(ctx, "SELECT * FROM users WHERE email = ?", email)

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

func (repo *UserRepository) GetUsersByMerchantID(merchantID int64) ([]*entity.User, error) {
	ctx := context.TODO()
	rows, err := repo.DB.QueryContext(ctx, "SELECT * FROM users WHERE merchant_id = ?", merchantID)
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

func (repo *UserRepository) CreateUser(param entity.CreateUserParam) (*entity.User, error) {
	ctx := context.TODO()
	res, err := repo.DB.ExecContext(
		ctx,
		"INSERT INTO users(name, email, role, position, password, merchant_id, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?)",
		param.Name,
		param.Email,
		param.Role,
		param.Position,
		param.Password,
		param.MerchantID,
		param.CreatedAt,
		param.CreatedAt)
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

func (repo *UserRepository) UpdateByID(userID int64, param entity.UpdateUserParam) error {
	ctx := context.TODO()
	res, err := repo.DB.ExecContext(
		ctx,
		"UPDATE users SET name = ?, email = ?, role = ?, position = ?, password = ?, updated_at = ? WHERE id = ?",
		param.Name, param.Email, param.Role, param.Position, param.Password, param.UpdatedAt, userID)
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

func (repo *UserRepository) DeleteUserByID(userID int64) error {
	ctx := context.TODO()
	res, err := repo.DB.ExecContext(ctx, "DELETE FROM users WHERE id = ?", userID)
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
