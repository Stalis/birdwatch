package stat

import "github.com/Stalis/birdwatch/pkg/stat/cpu"

type Info struct {
	CPULoad int
}

func GetInfo() (*Info, error) {
	return &Info{
		CPULoad: cpu.Load(),
	}, nil
}
