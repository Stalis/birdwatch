package api

import (
	"context"
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
	avgInterval := time.Second * time.Duration(req.Query.AveragingInterval)
	sendInterval := time.Second * time.Duration(req.Query.SendingInterval)

	watcher := mem.NewWatcher(avgInterval, sendInterval)
	watcher.Start()

	data := make(chan *mem.MemoryStat)
	errChan := make(chan error)

	go func(data <-chan *mem.MemoryStat, errChan chan<- error) {
		for {
			stat := <-data
			err := srv.Send(&pb.CurrentMemoryResponse{
				Total:     int32(stat.Total),
				Available: int32(stat.Available),
				Used:      int32(stat.Used),
			})
			if err != nil {
				errChan <- err
				break
			}
		}
	}(data, errChan)

	for {
		stat := watcher.Avg()

		select {
		case <-srv.Context().Done():
			return srv.Context().Err()
		case err := <-errChan:
			return err
		case data <- stat:
		}
	}
}