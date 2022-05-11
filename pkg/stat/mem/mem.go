package mem

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Stalis/birdwatch/pkg/config"
	"github.com/Stalis/birdwatch/pkg/utils"
	"go.uber.org/zap"
)

type MemoryStat struct {
	Available int
	Total     int
	Used      int

	LinuxVMStat  *LinuxVMStat
	DarwinVMStat *DarwinVMStat
}

func (m *MemoryStat) String() string {
	res, _ := json.MarshalIndent(m, "", "  ")
	return string(res)
}

type DarwinVMStat struct {
	Total          int
	PagesFree      int
	PagesActive    int
	PagesInactive  int
	PagesWiredDown int
}

type LinuxVMStat struct {
	MemTotal     int
	MemFree      int
	MemAvailable int
}

type Watcher struct {
	buffer            *utils.CircleBuffer
	averagingInterval time.Duration
	scanInterval      time.Duration
	cancelFunc        context.CancelFunc
}

func NewWatcher(averagingInterval time.Duration) *Watcher {
	zap.L().Debug("Creating new memory watcher",
		zap.Duration("averagingInterval", averagingInterval))

	cfg, _ := config.Get()
	scanInterval := cfg.Memory.ScanInterval

	bufferLength := averagingInterval.Milliseconds() / scanInterval.Milliseconds()
	zap.L().Debug("Buffer length", zap.Int64("bufferLength", bufferLength))

	return &Watcher{
		buffer:            utils.NewCircleBuffer(int(bufferLength)),
		averagingInterval: averagingInterval,
		scanInterval:      scanInterval,
	}
}

func (w *Watcher) Start(ctx context.Context) bool {
	if w.cancelFunc != nil {
		return false
	}

	zap.L().Debug("Start memory watcher")

	newCtx, cancelFunc := context.WithCancel(ctx)
	w.cancelFunc = cancelFunc

	dataCh := make(chan *MemoryStat)

	go func(ctx context.Context, dataCh chan<- *MemoryStat) {
		scanTicker := time.NewTicker(w.scanInterval)
		defer scanTicker.Stop()

		for {
			select {
			case <-ctx.Done():
				zap.L().Debug("Stop memory stats reading")
				return
			case <-scanTicker.C:
				zap.L().Debug("Read memory stats")
				dataCh <- GetMemoryStat(ctx)
			}
		}
	}(newCtx, dataCh)

	go func(ctx context.Context, dataCh <-chan *MemoryStat) {
		for {
			select {
			case <-ctx.Done():
				zap.L().Debug("Stop memory stats writing")
				return
			case v := <-dataCh:
				zap.L().Debug("Write memory stats")
				w.buffer.Add(v)
			}
		}
	}(newCtx, dataCh)

	return true
}

func (w *Watcher) Avg(ctx context.Context) *MemoryStat {
	zap.L().Debug("Counting average")

	closed := make(chan struct{})
	go func() {
		for {
			select {
			case <-ctx.Done():
				break
			default:
			}

			if w.buffer.Closed() {
				break
			}
		}
		closed <- struct{}{}
	}()

	<-closed
	close(closed)

	items := w.buffer.Items()
	res := &MemoryStat{}

	for _, v := range items {
		item := v.(*MemoryStat)
		res.Available += item.Available
		res.Used += item.Used
		res.Total += item.Total
	}

	res.Available /= len(items)
	res.Used /= len(items)
	res.Total /= len(items)

	zap.L().Debug("avg computed", zap.Stringer("avg", res))

	return res
}

func (w *Watcher) Stop() {
	zap.L().Debug("Stop memory watcher")
	w.cancelFunc()
}
