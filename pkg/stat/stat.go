package stat

import (
	"context"

	"github.com/Stalis/birdwatch/pkg/stat/cpu"
	"github.com/Stalis/birdwatch/pkg/stat/mem"
)

type Info struct {
	CPULoad int
	Memory  *mem.VirtualMemoryStat
}

func GetInfo() (*Info, error) {
	return &Info{
		CPULoad: cpu.Load(),
		Memory:  mem.GetVirtualMemoryStat(context.TODO()),
	}, nil
}
