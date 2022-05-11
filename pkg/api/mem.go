package api

import (
	"context"
	"time"

	"github.com/Stalis/birdwatch/pkg/api/pb"
	"github.com/Stalis/birdwatch/pkg/stat/mem"
	"go.uber.org/zap"
)

type MemoryServer struct {
	pb.UnimplementedMemoryServer
}

func NewMemoryServer() *MemoryServer {
	return &MemoryServer{}
}

func (m *MemoryServer) GetCurrentMemoryStats(ctx context.Context, req *pb.CurrentMemoryRequest) (*pb.CurrentMemoryResponse, error) { //nolint:lll
	stat := mem.GetMemoryStat(ctx)

	return convertMemoryStat(stat), nil
}

func (m *MemoryServer) GetMemoryStats(req *pb.MemoryStatsRequest, srv pb.Memory_GetMemoryStatsServer) error {
	zap.L().Debug("Request GetMemoryStats")
	defer zap.L().Debug("End of GetMemoryStats request")

	avgInterval := time.Second * time.Duration(req.Query.AveragingInterval)
	sendInterval := time.Second * time.Duration(req.Query.SendingInterval)

	watcher := mem.NewWatcher(avgInterval)
	watcher.Start(srv.Context())
	defer watcher.Stop()

	data := make(chan *mem.MemoryStat)
	errChan := make(chan error)

	go func(data <-chan *mem.MemoryStat, errChan chan<- error) {
		for {
			stat := <-data
			response := convertMemoryStat(stat)
			zap.L().Debug("Send memory stats", zap.Stringer("data", response))
			err := srv.Send(response)
			if err != nil {
				errChan <- err
				break
			}
		}
	}(data, errChan)

	sendTicker := time.NewTicker(sendInterval)
	defer sendTicker.Stop()

	for {
		stat := watcher.Avg(srv.Context())

		select {
		case <-srv.Context().Done():
			return srv.Context().Err()
		case err := <-errChan:
			zap.L().Error("Error while sending memory data", zap.Error(err))
		case data <- stat:
			zap.L().Debug("Send memory stats", zap.Stringer("data", stat))
		}

		<-sendTicker.C
	}
}

func convertMemoryStat(stat *mem.MemoryStat) *pb.CurrentMemoryResponse {
	return &pb.CurrentMemoryResponse{
		Total:     int64(stat.Total),
		Available: int64(stat.Available),
		Used:      int64(stat.Used),
	}
}
