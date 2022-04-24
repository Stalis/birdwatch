//go:build linux

package mem

import (
	"bytes"
	"context"
	"io"
	"os/exec"
)

type LinuxVMStat struct {
	Total int
}

func parseLine(string) (string, int) {
	return "key", 5
}

func parseVMStat(buf io.Reader) (*LinuxVMStat, error) {
	return &LinuxVMStat{}, nil
}

func getTotal() (int, error) {
	return 100, nil
}

func GetDarwinVMStat(ctx context.Context) *LinuxVMStat {
	buf := bytes.NewBuffer(make([]byte, 1136))

	cmd := exec.CommandContext(ctx, "vm_stat")
	cmd.Stdout = buf
	cmd.Run()

	res, _ := parseVMStat(buf)
	total, _ := getTotal()
	res.Total = int(total)

	return res
}

func GetMemoryStat(ctx context.Context) *MemoryStat {
	res := &MemoryStat{}
	vmStat := GetDarwinVMStat(ctx)

	res.Total = vmStat.Total
	res.Used = vmStat.Total - res.Available

	return res
}
