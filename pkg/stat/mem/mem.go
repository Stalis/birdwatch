package mem

import (
	"context"
	"time"

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
}

func NewWatcher(averagingInterval time.Duration, scanInterval time.Duration) *Watcher {
	return &Watcher{
		buffer:            utils.NewCircleBuffer(int(averagingInterval / scanInterval)),
		averagingInterval: averagingInterval,
		scanInterval:      scanInterval,
	}
}

func (w *Watcher) Start() {
}

func (w *Watcher) Avg() *MemoryStat {
	return GetMemoryStat(context.TODO())
}

func (w *Watcher) Done() {
}
