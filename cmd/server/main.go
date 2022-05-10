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

const (
	CodeOk = iota
	CodeConfigInit
	CodeLoggerInit
	CodeTCPListener
	CodeGrpcServer
)

func main() {
	code := server()
	os.Exit(code)
}

func server() int {
	cfg, err := config.InitConfig()
	if err != nil {
		fmt.Println(err)
		return CodeConfigInit
	}

	logger, err := log.InitZapLogger(&log.Config{
		Console: cfg.Logging.Verbose,
		Level:   cfg.Logging.Level,
		File:    cfg.Logging.File,
	})
	if err != nil {
		fmt.Println(err)
		return CodeLoggerInit
	}
	defer logger.Sync()
	logger.Info("Initialize logger")

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		zap.S().Errorf("Error while setup tcp listener: %v", err)
		return CodeTCPListener
	}
	zap.S().Info("Started tcp listener")

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	prepareGrpcServer(grpcServer)
	zap.S().Info("Initialized gRPC server")

	if err := grpcServer.Serve(lis); err != nil {
		zap.S().Errorf("Error while serving gRPC server: %v", err)
		return CodeGrpcServer
	}

	return CodeOk
}

func prepareGrpcServer(srv *grpc.Server) {
	memoryServer := api.NewMemoryServer()
	pb.RegisterMemoryServer(srv, memoryServer)
}
