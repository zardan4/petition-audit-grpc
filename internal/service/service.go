package service

import (
	"context"

	storage "github.com/zardan4/petition-audit-grpc/internal/storage/mongo"
	"github.com/zardan4/petition-audit-grpc/pkg/core/audit"
)

type Audit interface {
	Log(ctx context.Context, audit *audit.LogRequest) error
}

type Service struct {
	Audit
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		Audit: NewAuditService(storage),
	}
}
