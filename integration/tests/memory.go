package tests

import (
	"context"
	"time"

	"github.com/Stalis/birdwatch/pkg/api/pb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	AvgInterval  = 2 * time.Second
	SendInterval = 2 * time.Second
)

const (
	ErrInvalidRequest = "invalid request"
	ErrReceiveFailed  = "receive failed"
)

func TestMemoryStats(baseCtx context.Context, conn *grpc.ClientConn) error {
	client := pb.NewMemoryClient(conn)

	ctx, cancel := context.WithCancel(baseCtx)
	defer cancel()

	stream, err := client.GetMemoryStats(ctx, &pb.MemoryStatsRequest{
		Query: &pb.Query{
			AveragingInterval: int32(AvgInterval.Seconds()),
			SendingInterval:   int32(SendInterval.Seconds()),
		},
	})
	if err != nil {
		return errors.Wrap(err, ErrInvalidRequest)
	}

	for i := 0; i < 10; i++ {
		_, err := stream.Recv()
		if err != nil {
			errors.Wrap(err, ErrReceiveFailed)
		}
	}

	return nil
}
