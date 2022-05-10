//go:build linux
// +build linux

package mem

import (
	"bufio"
	"context"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func parseLine(input string) (string, int) {
	parts := strings.Split(input, ":")
	if len(parts) <= 1 {
		return "", 0
	}

	key := parts[0]

	value := parts[1]
	value = strings.TrimFunc(value, func(r rune) bool {
		return !unicode.IsDigit(r)
	})
	v, _ := strconv.Atoi(value)

	return key, v
}

func parseVMStat(input io.Reader) (*LinuxVMStat, error) {
	s := bufio.NewScanner(input)
	res := &LinuxVMStat{}

	for s.Scan() {
		key, value := parseLine(s.Text())

		switch key {
		case "MemTotal":
			res.MemTotal = value
		case "MemAvailable":
			res.MemAvailable = value
		case "MemFree":
			res.MemFree = value
		}
	}

	return res, nil
}

func GetLinuxVMStat(ctx context.Context) *LinuxVMStat {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		zap.S().Errorf("Unable to read `/proc/meminfo`: %v", err)
		return nil
	}
	defer file.Close()

	res, _ := parseVMStat(file)
	return res
}

func GetMemoryStat(ctx context.Context) *MemoryStat {
	res := &MemoryStat{}

	if vmStat := GetLinuxVMStat(ctx); vmStat != nil {
		res.LinuxVMStat = vmStat
		res.Available = vmStat.MemAvailable
		res.Total = vmStat.MemTotal
		res.Used = vmStat.MemTotal - res.Available
	}

	return res
}
