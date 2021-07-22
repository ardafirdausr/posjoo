package entity

import "time"

type Product struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	PhotoUrl   *string   `json:"photo_url"`
	SKU        string    `json:"sku"`
	Quantity   int       `json:"quantity"`
	MerchantID int64     `json:"merchant_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateProductParam struct {
	Name       string    `json:"name"`
	SKU        string    `json:"sku"`
	MerchantID int64     `json:"-"`
	CreatedAt  time.Time `json:"-"`
}

type UpdatedProductparam struct {
	Name      string    `json:"name"`
	SKU       string    `json:"sku"`
	UpdatedAt time.Time `json:"-"`
}
