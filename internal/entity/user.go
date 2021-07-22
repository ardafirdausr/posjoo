package entity

import "time"

type UserRole string

var (
	UserRoleOwner   UserRole = "user.role.owner"
	UserRoleManager UserRole = "user.role.manager"
	UserRoleStaff   UserRole = "user.role.staff"
)

type User struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	PhotoUrl   *string   `json:"photo_url"`
	Role       UserRole  `json:"role"`
	Position   string    `json:"position"`
	Password   string    `json:"password"`
	MerchantID int64     `json:"merchant_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateUserParam struct {
	Name                 string   `json:"name" validate:"required"`
	Email                string   `json:"email" validate:"required,email"`
	Role                 UserRole `json:"role" validate:"required"`
	Position             string   `json:"position" validate:"required"`
	Password             string   `json:"password" validate:"required"`
	PasswordConfirmation string   `json:"password_confirmation" validate:"required,eqfield=Password"`
	MerchantID           int64    `json:"-"`
}

type UpdateUserParam struct {
	Name                 string   `json:"name" validate:"required"`
	Email                string   `json:"email" validate:"required,email"`
	Role                 UserRole `json:"role" validate:"required"`
	Position             string   `json:"position" validate:"required"`
	Password             string   `json:"password" validate:"required"`
	PasswordConfirmation string   `json:"password_confirmation" validate:"required,eqfield=Password"`
}

type UpdateUserPhotoParam struct {
	PhotoUrl *string `json:"photo_url" validate:"required"`
}
