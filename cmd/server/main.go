package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Stalis/birdwatch/pkg/api"
	"github.com/Stalis/birdwatch/pkg/api/pb"
	"github.com/Stalis/birdwatch/pkg/config"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatal(err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	prepareGrpcServer(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func prepareGrpcServer(srv *grpc.Server) {
	memoryServer := api.NewMemoryServer()
	pb.RegisterMemoryServer(srv, memoryServer)
}
