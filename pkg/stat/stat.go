package stat

import (
	"context"
	"encoding/json"

	"github.com/Stalis/birdwatch/pkg/stat/cpu"
	"github.com/Stalis/birdwatch/pkg/stat/mem"
)

type Info struct {
	CPULoad int
	Memory  *mem.MemoryStat
}

func (i *Info) String() string {
	res, _ := json.MarshalIndent(i, "", "  ")
	return string(res)
}

func GetInfo() (*Info, error) {
	return &Info{
		CPULoad: cpu.Load(),
		Memory:  mem.GetMemoryStat(context.TODO()),
	}, nil
}
