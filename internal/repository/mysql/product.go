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

func (repo *ProductRepository) GetProductByID(productID int64) (*entity.Product, error) {
	ctx := context.TODO()
	row := repo.DB.QueryRowContext(ctx, "SELECT * FROM products WHERE id = ?", productID)

	var product entity.Product
	var err = row.Scan(
		&product.ID,
		&product.Name,
		&product.PhotoUrl,
		&product.SKU,
		&product.MerchantID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		err := entity.ErrNotFound{
			Message: "Product not found",
			Err:     err,
		}
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *ProductRepository) GetProductBySKU(SKU string) (*entity.Product, error) {
	ctx := context.TODO()
	row := repo.DB.QueryRowContext(ctx, "SELECT * FROM products WHERE sku = ?", SKU)

	var product entity.Product
	var err = row.Scan(
		&product.ID,
		&product.Name,
		&product.PhotoUrl,
		&product.SKU,
		&product.MerchantID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		err := entity.ErrNotFound{
			Message: "Product not found",
			Err:     err,
		}
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *ProductRepository) GetProductsByMerchantID(merchantID int64) ([]*entity.Product, error) {
	ctx := context.TODO()
	rows, err := repo.DB.QueryContext(ctx, "SELECT * FROM products WHERE merchant_id = ?", merchantID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	rows.Close()

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

func (repo *ProductRepository) CreateProduct(param entity.CreateProductParam) (*entity.Product, error) {
	ctx := context.TODO()
	res, err := repo.DB.ExecContext(
		ctx,
		"INSERT INTO products(name, sku, merchant_id, created_at, updated_at)",
		param.Name, param.SKU, param.MerchantID, param.CreatedAt, param.CreatedAt)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	ID, err := res.LastInsertId()
	if err != nil {
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

func (repo *ProductRepository) UpdateProductByID(productId int64, param entity.UpdatedProductparam) error {
	ctx := context.TODO()
	res, err := repo.DB.ExecContext(
		ctx,
		"UPDATE products SET name = ?, sku = ?, updated_at = ? WHERE id = ?",
		param.Name, param.SKU, param.UpdatedAt, productId)
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

func (repo *ProductRepository) DeleteProductByID(productId int64) error {
	ctx := context.TODO()
	res, err := repo.DB.ExecContext(ctx, "DELETE FROM products WHERE id = ?", productId)
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
