//go:build darwin
// +build darwin

package cpu

import (
	"golang.org/x/sys/unix"
)

func Load() int {
	count, _ := unix.SysctlUint32("machdep.cpu.core_count")
	return int(count)
}

func GetStat() *Stat {
	return &Stat{
		Load: Load(),
	}
}
