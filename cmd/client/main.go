package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Stalis/birdwatch/pkg/api/pb"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	Host         string
	Port         int32
	AvgInterval  time.Duration
	SendInterval time.Duration
}

func main() {
	var opts []grpc.DialOption

	config, err := getConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", config.Host, config.Port), opts...)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	listenMemoryStats(context.Background(), conn, config)

	fmt.Println(config)
}

func getConfig() (*Config, error) {
	f := pflag.NewFlagSet("config", pflag.ExitOnError)
	f.Usage = func() {
		fmt.Println(f.FlagUsages())
		os.Exit(0)
	}

	f.StringP("host", "h", "localhost", "hostname of birdwatch server")
	f.Int32P("port", "p", 50051, "port of birdwatch server listening")
	f.DurationP("avg-interval", "a", time.Second*2, "averaging stats interval")
	f.DurationP("send-interval", "s", time.Second, "send stats interval")

	if err := f.Parse(os.Args[1:]); err != nil {
		return nil, err
	}

	host, err := f.GetString("host")
	if err != nil {
		return nil, err
	}

	port, err := f.GetInt32("port")
	if err != nil {
		return nil, err
	}

	avgInterval, err := f.GetDuration("avg-interval")
	if err != nil {
		return nil, err
	}
	avgInterval = avgInterval.Round(time.Second)

	sendInterval, err := f.GetDuration("send-interval")
	if err != nil {
		return nil, err
	}
	sendInterval = sendInterval.Round(time.Second)

	res := &Config{
		Host:         host,
		Port:         port,
		AvgInterval:  avgInterval,
		SendInterval: sendInterval,
	}

	return res, nil
}

func listenMemoryStats(baseCtx context.Context, conn *grpc.ClientConn, config *Config) {
	client := pb.NewMemoryClient(conn)

	ctx, cancel := context.WithCancel(baseCtx)
	defer cancel()

	stream, err := client.GetMemoryStats(ctx, &pb.MemoryStatsRequest{
		Query: &pb.Query{
			AveragingInterval: int32(config.AvgInterval.Seconds()),
			SendingInterval:   int32(config.SendInterval.Seconds()),
		},
	})
	if err != nil {
		fmt.Printf("Error while request memory stats stream: %v\n", err)
	}

	fmt.Printf("Available\tTotal\tUsed\n")
	for {
		data, err := stream.Recv()
		if err != nil {
			fmt.Printf("Receive memory stats failed: %v\n", err)
			break
		}
		fmt.Printf("%v\t%v\t%v\n", data.Available, data.Total, data.Used)
	}
}
