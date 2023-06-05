package storage

import (
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/config"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/pkg/db"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/pkg/logger"
)

type StorageI interface {
}

type StoragePg struct {
}

// NewStoragePg
func New(db *db.Postgres, log *logger.Logger, cfg *config.Config) StorageI {
	return &StoragePg{}
}