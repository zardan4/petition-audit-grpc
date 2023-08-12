package service

import (
	"context"

	storage "github.com/zardan4/petition-audit-grpc/internal/storage/mongo"
	"github.com/zardan4/petition-audit-grpc/pkg/core/audit"
)

type AuditService struct {
	storage *storage.Storage
}

func NewAuditService(storage *storage.Storage) *AuditService {
	return &AuditService{
		storage: storage,
	}
}

func (s *AuditService) Log(ctx context.Context, log *audit.LogRequest) error {
	item := audit.LogItem{
		Entity:    log.GetEntity().String(),
		Action:    log.GetAction().String(),
		EntityID:  log.GetEntityId(),
		Timestamp: log.GetTimestamp().AsTime(),
	}

	return s.storage.Insert(ctx, item)
}
