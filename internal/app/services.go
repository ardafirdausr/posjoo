package app

import (
	"os"

	"github.com/ardafirdausr/posjoo-server/internal"
	"github.com/ardafirdausr/posjoo-server/internal/pkg/storage"
)

type services struct {
	storageService internal.Storage
}

func newServices() *services {
	services := new(services)
	services.storageService = storage.NewFileSystemStorage("storage", os.Getenv("APP_DOMAIN"))
	return services
}
