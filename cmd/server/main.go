package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/Stalis/birdwatch/pkg/api"
	"github.com/Stalis/birdwatch/pkg/api/pb"
	"google.golang.org/grpc"
)

var (
	host string
	port int
)

func main() {
	flag.IntVar(&port, "port", 50051, "Port for GRPC server")
	flag.StringVar(&host, "host", "localhost", "Host for GRPC server")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
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
