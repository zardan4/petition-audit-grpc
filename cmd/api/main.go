package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/zardan4/petition-audit-grpc/internal/config"
	server "github.com/zardan4/petition-audit-grpc/internal/server/grpc"
	"github.com/zardan4/petition-audit-grpc/internal/service"
	storage "github.com/zardan4/petition-audit-grpc/internal/storage/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		logrus.Fatal(err)
	}

	// init db connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client()
	opts.ApplyURI(cfg.DB.ConnectionLine)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		logrus.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logrus.Fatal(err)
	}

	db := client.Database(cfg.DB.Database)
	storage := storage.NewStorage(db)
	service := service.NewService(storage)

	auditSrv := server.NewAuditServer(service)
	srv := server.NewServer(auditSrv)

	fmt.Printf("Server started at %s", time.Now())

	if err := srv.ListenAndServe(cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
