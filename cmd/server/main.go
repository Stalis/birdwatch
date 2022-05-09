package main

import (
	"fmt"
	"net"
	"os"

	"github.com/Stalis/birdwatch/pkg/api"
	"github.com/Stalis/birdwatch/pkg/api/pb"
	"github.com/Stalis/birdwatch/pkg/config"
	"github.com/Stalis/birdwatch/pkg/log"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		os.Exit(1)
	}

	logger, err := log.InitZapLogger(&log.Config{
		Console: cfg.Logging.Verbose,
		Level:   cfg.Logging.Level,
		File:    cfg.Logging.File,
	})
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		os.Exit(2)
	}
	defer logger.Sync()
	logger.Info("Initialize logger")

	sugared := logger.Sugar()
	sugared.Debug("Initialize logger")

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		sugared.Errorf("Error while setup tcp listener: %v", err)
		return
	}
	sugared.Info("Started tcp listener")

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	prepareGrpcServer(grpcServer)
	sugared.Info("Initialized gRPC server")

	if err := grpcServer.Serve(lis); err != nil {
		sugared.Errorf("Error while serving gRPC server: %v", err)
	}
}

func prepareGrpcServer(srv *grpc.Server) {
	memoryServer := api.NewMemoryServer()
	pb.RegisterMemoryServer(srv, memoryServer)
}
