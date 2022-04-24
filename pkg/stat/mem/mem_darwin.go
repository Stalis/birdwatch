//go:build darwin
// +build darwin

package mem

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"os/exec"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

func parseVMStat(input io.Reader) (*DarwinVMStat, error) {
	s := bufio.NewScanner(input)
	res := &DarwinVMStat{}

	for s.Scan() {
		key, value := parseLine(s.Text())

		switch key {
		case "Pages free":
			res.PagesFree = value
		case "Pages active":
			res.PagesActive = value
		case "Pages inactive":
			res.PagesInactive = value
		case "Pages wired down":
			res.PagesWiredDown = value
		}
	}

	return res, nil
}

func getTotal() (uint64, error) {
	total, err := unix.SysctlUint64("hw.memsize")
	if err != nil {
		return 0, err
	}
	return total, nil
}

func parseLine(input string) (string, int) {
	parts := strings.Split(input, ":")
	if len(parts) <= 1 {
		return "", 0
	}
	key := parts[0]
	key = strings.Trim(key, "\"")

	value := parts[1]
	value = strings.Trim(value, " .")

	v, _ := strconv.Atoi(value)
	return key, v
}

func GetDarwinVMStat(ctx context.Context) *DarwinVMStat {
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

	res.DarwinVMStat = vmStat
	res.Available = vmStat.PagesFree + vmStat.PagesInactive
	res.Total = vmStat.Total
	res.Used = vmStat.Total - res.Available

	return res
}
