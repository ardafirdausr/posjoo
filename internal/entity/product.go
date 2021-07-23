package entity

import "time"

type Product struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	SKU        string    `json:"sku"`
	PhotoUrl   *string   `json:"photo_url"`
	MerchantID int64     `json:"merchant_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateProductParam struct {
	Name       string    `json:"name" validate:"required"`
	SKU        string    `json:"sku" validate:"required"`
	MerchantID int64     `json:"-"`
	CreatedAt  time.Time `json:"-"`
}

type UpdatedProductparam struct {
	Name      string    `json:"name" validate:"required"`
	SKU       string    `json:"sku" validate:"required"`
	UpdatedAt time.Time `json:"-"`
}

type UpdateProductPhotoParam struct {
	PhotoUrl *string `json:"photo_url" validate:"required"`
}
