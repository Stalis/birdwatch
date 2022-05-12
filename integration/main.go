package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Stalis/birdwatch/integration/tests"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getTarget() (string, error) {
	f := pflag.NewFlagSet("config", pflag.ExitOnError)
	f.Usage = func() {
		fmt.Println(f.FlagUsages())
		os.Exit(0)
	}

	f.StringP("host", "h", "localhost", "hostname of birdwatch server")
	f.Int32P("port", "p", 50051, "port of birdwatch server listening")

	if err := f.Parse(os.Args[1:]); err != nil {
		return "", err
	}

	host, err := f.GetString("host")
	if err != nil {
		return "", err
	}

	port, err := f.GetInt32("port")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v:%v", host, port), nil
}

func main() {
	var opts []grpc.DialOption

	target, err := getTarget()
	if err != nil {
		log.Println(err)
		return
	}

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	if err := tests.TestMemoryStats(context.Background(), conn); err != nil {
		log.Println("FAIL!")
		log.Println(err)
		return
	}

	log.Println("PASS!")
}
