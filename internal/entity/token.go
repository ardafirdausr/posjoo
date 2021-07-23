package entity

import "github.com/dgrijalva/jwt-go"

type TokenPayload struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	PhotoUrl string `json:"photo_url"`
}

type JWTPayload struct {
	TokenPayload
	jwt.StandardClaims
}

type GoogleAuth struct {
	TokenID  string `json:"token_id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	PhotoUrl string `json:"photo_url,omitempty"`
}
