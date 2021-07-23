package internal

import "github.com/ardafirdausr/posjoo-server/internal/entity"

type Tokenizer interface {
	Generate(entity.TokenPayload) (string, error)
	Parse(token string) (*entity.TokenPayload, error)
}
