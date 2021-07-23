package mysql

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/ardafirdausr/posjoo-server/internal/entity"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(DB *sql.DB) *ProductRepository {
	repo := new(ProductRepository)
	repo.DB = DB
	return repo
}

func (repo ProductRepository) GetProductByID(ctx context.Context, productID int64) (*entity.Product, error) {
	var query = "SELECT * FROM products WHERE id = ?"
	var row *sql.Row
	txKey := transactionContextKey("tx")
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		row = tx.QueryRow(query, productID)
	} else {
		row = repo.DB.QueryRowContext(ctx, query, productID)
	}

	var product entity.Product
	var err = row.Scan(
		&product.ID,
		&product.Name,
		&product.SKU,
		&product.PhotoUrl,
		&product.MerchantID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		errNotFound := entity.ErrNotFound{
			Message: "Product not found",
			Err:     err,
		}
		return nil, errNotFound
	}

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo ProductRepository) GetProductBySKU(ctx context.Context, SKU string) (*entity.Product, error) {
	var query = "SELECT * FROM products WHERE sku = ?"
	var row *sql.Row
	txKey := transactionContextKey("tx")
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		row = tx.QueryRow(query, SKU)
	} else {
		row = repo.DB.QueryRowContext(ctx, query, SKU)
	}

	var product entity.Product
	var err = row.Scan(
		&product.ID,
		&product.Name,
		&product.SKU,
		&product.PhotoUrl,
		&product.MerchantID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		errNotFound := entity.ErrNotFound{
			Message: "Product not found",
			Err:     err,
		}
		return nil, errNotFound
	}

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo ProductRepository) GetProductsByMerchantID(ctx context.Context, merchantID int64) ([]*entity.Product, error) {
	var query = "SELECT * FROM products WHERE merchant_id = ?"
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
	defer rows.Close()

	products := []*entity.Product{}
	for rows.Next() {
		var product entity.Product
		var err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.SKU,
			&product.PhotoUrl,
			&product.MerchantID,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		products = append(products, &product)
	}
	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return products, nil
}

func (repo ProductRepository) CreateProduct(ctx context.Context, param entity.CreateProductParam) (*entity.Product, error) {
	var query = "INSERT INTO products(name, sku, merchant_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	var res sql.Result
	var err error
	txKey := transactionContextKey("tx")
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		res, err = tx.Exec(query, param.Name, param.SKU, param.MerchantID, param.CreatedAt, param.CreatedAt)
	} else {
		res, err = repo.DB.ExecContext(ctx, query, param.Name, param.SKU, param.MerchantID, param.CreatedAt, param.CreatedAt)
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

	product := &entity.Product{
		ID:         ID,
		Name:       param.Name,
		SKU:        param.SKU,
		MerchantID: param.MerchantID,
		CreatedAt:  param.CreatedAt,
		UpdatedAt:  param.CreatedAt,
	}

	return product, nil
}

func (repo ProductRepository) UpdateProductByID(ctx context.Context, productId int64, param entity.UpdatedProductparam) error {
	var query = "UPDATE products SET name = ?, sku = ?, updated_at = ? WHERE id = ?"
	var res sql.Result
	var err error
	txKey := transactionContextKey("tx")
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		res, err = tx.Exec(query, param.Name, param.SKU, param.UpdatedAt, productId)
	} else {
		res, err = repo.DB.ExecContext(ctx, query, param.Name, param.SKU, param.UpdatedAt, productId)
	}

	if err != nil {
		log.Println(err.Error())
		return err
	}

	if num, _ := res.RowsAffected(); num < 1 {
		err := errors.New("failed to update product")
		enf := entity.ErrNotFound{
			Message: "Product not found",
			Err:     err,
		}
		log.Println(enf.Error())
		return enf
	}

	return nil
}

func (repo ProductRepository) DeleteProductByID(ctx context.Context, productId int64) error {
	var query = "DELETE FROM products WHERE id = ?"
	var res sql.Result
	var err error
	txKey := transactionContextKey("tx")
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		res, err = tx.Exec(query, productId)
	} else {
		res, err = repo.DB.ExecContext(ctx, query, productId)
	}

	if err != nil {
		log.Println(err.Error())
		return err
	}

	if num, _ := res.RowsAffected(); num < 1 {
		err := errors.New("failed to delete product")
		enf := entity.ErrNotFound{
			Message: "Product not found",
			Err:     err,
		}
		log.Println(enf.Error())
		return enf
	}

	return nil
}
