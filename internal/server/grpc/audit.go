package server

import (
	"context"

	"github.com/zardan4/petition-audit-grpc/internal/service"
	"github.com/zardan4/petition-audit-grpc/pkg/core/audit"
)

type AuditServer struct {
	audit.AuditServiceServer

	service *service.Service
}

func NewAuditServer(service *service.Service) *AuditServer {
	return &AuditServer{
		service: service,
	}
}

func (s *AuditServer) Log(ctx context.Context, log *audit.LogRequest) (*audit.Empty, error) {
	return &audit.Empty{}, s.service.Log(ctx, log)
}
