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
	cfg, err := config.Get()
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
	defer func() {
		if msg := recover(); msg != nil {
			logger.Error("Server panic", zap.Any("panicError", msg))
		}
	}()

	logger.Info("Initialize logger")

	target := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	lis, err := net.Listen("tcp", target)
	if err != nil {
		logger.Error("Error while setup tcp listener", zap.Error(err))
		return CodeTCPListener
	}
	logger.Info("Started tcp listener", zap.String("target", target))

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	prepareGrpcServer(grpcServer)
	logger.Info("Initialized gRPC server")

	if err := grpcServer.Serve(lis); err != nil {
		logger.Error("Error while serving gRPC server", zap.Error(err))
		return CodeGrpcServer
	}

	logger.Info("Server stopped successful")
	return CodeOk
}

func prepareGrpcServer(srv *grpc.Server) {
	memoryServer := api.NewMemoryServer()
	pb.RegisterMemoryServer(srv, memoryServer)
}
