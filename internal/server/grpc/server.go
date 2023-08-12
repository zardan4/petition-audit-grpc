package server

import (
	"fmt"
	"net"

	"github.com/zardan4/petition-audit-grpc/pkg/core/audit"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server

	auditServer audit.AuditServiceServer
}

func NewServer(auditServer audit.AuditServiceServer) *Server {
	return &Server{
		grpcServer:  grpc.NewServer(),
		auditServer: auditServer,
	}
}

func (s *Server) ListenAndServe(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	audit.RegisterAuditServiceServer(s.grpcServer, s.auditServer)

	if err := grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}
