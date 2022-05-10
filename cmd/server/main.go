package main

import (
	"fmt"
	"net"
	"os"

	"github.com/Stalis/birdwatch/pkg/api"
	"github.com/Stalis/birdwatch/pkg/api/pb"
	"github.com/Stalis/birdwatch/pkg/config"
	"github.com/Stalis/birdwatch/pkg/log"
	"go.uber.org/zap"
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

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		zap.S().Errorf("Error while setup tcp listener: %v", err)
		return
	}
	zap.S().Info("Started tcp listener")

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	prepareGrpcServer(grpcServer)
	zap.S().Info("Initialized gRPC server")

	if err := grpcServer.Serve(lis); err != nil {
		zap.S().Errorf("Error while serving gRPC server: %v", err)
	}
}

func prepareGrpcServer(srv *grpc.Server) {
	memoryServer := api.NewMemoryServer()
	pb.RegisterMemoryServer(srv, memoryServer)
}
