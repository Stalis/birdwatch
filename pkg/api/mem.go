package api

import (
	"context"
	"fmt"
	"time"

	"github.com/Stalis/birdwatch/pkg/api/pb"
	"github.com/Stalis/birdwatch/pkg/stat/mem"
)

type MemoryServer struct {
	pb.UnimplementedMemoryServer
}

func NewMemoryServer() *MemoryServer {
	return &MemoryServer{}
}

func (m *MemoryServer) GetCurrentMemoryStats(ctx context.Context, req *pb.CurrentMemoryRequest) (*pb.CurrentMemoryResponse, error) { //nolint:lll
	stat := mem.GetMemoryStat(ctx)

	return &pb.CurrentMemoryResponse{
		Total:     int32(stat.Total),
		Available: int32(stat.Available),
		Used:      int32(stat.Used),
	}, nil
}

func (m *MemoryServer) GetMemoryStats(req *pb.MemoryStatsRequest, srv pb.Memory_GetMemoryStatsServer) error {
	fmt.Println("Call GetMemoryStats")
	for {
		time.Sleep(time.Duration(req.Interval) * time.Millisecond)

		stat := mem.GetMemoryStat(srv.Context())
		err := srv.Send(&pb.CurrentMemoryResponse{
			Total:     int32(stat.Total),
			Available: int32(stat.Available),
			Used:      int32(stat.Used),
		})

		if err != nil {
			return err
		}
	}

	fmt.Println("Exit GetMemoryStats")
	return nil
}
