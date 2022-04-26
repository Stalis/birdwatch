package api

import (
	"context"

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
	for srv.Context()
	req.Interval
}
