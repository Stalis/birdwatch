package mem

import (
	"context"
	"time"

	"github.com/Stalis/birdwatch/pkg/config"
	"github.com/Stalis/birdwatch/pkg/utils"
)

type MemoryStat struct {
	Available int
	Total     int
	Used      int

	LinuxVMStat  *LinuxVMStat
	DarwinVMStat *DarwinVMStat
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

func NewWatcher(averagingInterval time.Duration, scanInterval time.Duration) *Watcher {
	cfg, _ := config.Get()

	return &Watcher{
		buffer:            utils.NewCircleBuffer(int(averagingInterval / scanInterval)),
		averagingInterval: averagingInterval,
		scanInterval:      cfg.Memory.ScanInterval,
	}
}

func (w *Watcher) Start(ctx context.Context) bool {
	if w.cancelFunc != nil {
		return false
	}

	newCtx, cancelFunc := context.WithCancel(ctx)
	w.cancelFunc = cancelFunc

	dataCh := make(chan *MemoryStat)

	go func(ctx context.Context, dataCh chan<- *MemoryStat) {
		scanTimer := time.NewTimer(w.scanInterval)

		for {
			select {
			case <-ctx.Done():
				return
			case <-scanTimer.C:
				dataCh <- GetMemoryStat(ctx)
			}
		}
	}(newCtx, dataCh)

	go func(ctx context.Context, dataCh <-chan *MemoryStat) {
		for {
			select {
			case <-ctx.Done():
				return
			case v := <-dataCh:
				w.buffer.Add(v)
			}
		}
	}(newCtx, dataCh)

	return true
}

func (w *Watcher) Avg(ctx context.Context) *MemoryStat {
	closed := make(chan struct{})
	go func() {
		for {
			if w.buffer.Closed() {
				closed <- struct{}{}
				return
			}

			select {
			case <-ctx.Done():
				return
			default:
			}
		}
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

	return res
}

func (w *Watcher) Done() {
	w.cancelFunc()
}
