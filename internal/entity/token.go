package entity

import "github.com/dgrijalva/jwt-go"

type TokenPayload struct {
	ID         int64    `json:"id"`
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	PhotoUrl   *string  `json:"photo_url"`
	Role       UserRole `json:"role"`
	Position   string   `json:"position"`
	MerchantID int64    `json:"merchant_id"`
}

type JWTPayload struct {
	TokenPayload
	jwt.StandardClaims
}
