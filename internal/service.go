package internal

import (
	"mime/multipart"

	"github.com/ardafirdausr/posjoo-server/internal/entity"
)

type Tokenizer interface {
	Generate(entity.TokenPayload) (string, error)
	Parse(token string) (*entity.TokenPayload, error)
}

type Storage interface {
	Save(file *multipart.FileHeader, dir string, filename string) (string, error)
}
