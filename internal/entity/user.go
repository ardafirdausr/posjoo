package entity

import "time"

type UserRole string

var (
	UserRoleOwner   UserRole = "owner"
	UserRoleManager UserRole = "manager"
	UserRoleStaff   UserRole = "staff"
)

type User struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	PhotoUrl   *string   `json:"photo_url"`
	Role       UserRole  `json:"role"`
	Position   string    `json:"position"`
	Password   string    `json:"-"`
	MerchantID int64     `json:"merchant_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type RegisterParam struct {
	Name                 string `json:"name" validate:"required"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
	BusinessName         string `json:"business_name" validate:"required"`
	BusinessAddress      string `json:"business_address" validate:"required"`
	BusinessPhone        string `json:"business_phone" validate:"required"`
}

type LoginParam struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type CreateUserParam struct {
	Name                 string    `json:"name" validate:"required"`
	Email                string    `json:"email" validate:"required,email"`
	Role                 UserRole  `json:"role" validate:"required"`
	Position             string    `json:"position" validate:"required"`
	Password             string    `json:"password" validate:"required"`
	PasswordConfirmation string    `json:"password_confirmation" validate:"required,eqfield=Password"`
	MerchantID           int64     `json:"-" validate:"required"`
	CreatedAt            time.Time `json:"-"`
}

type UpdateUserParam struct {
	Name                 string    `json:"name" validate:"required"`
	Email                string    `json:"email" validate:"required,email"`
	Role                 UserRole  `json:"role" validate:"required"`
	Position             string    `json:"position" validate:"required"`
	Password             string    `json:"password" validate:"required"`
	PasswordConfirmation string    `json:"password_confirmation" validate:"required,eqfield=Password"`
	UpdatedAt            time.Time `json:"-"`
}

type UpdateUserPhotoParam struct {
	PhotoUrl *string `json:"photo_url" validate:"required"`
}
