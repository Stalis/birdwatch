package mem

import (
	"bufio"
	"bytes"
	"context"
	"os/exec"
	"strconv"
	"strings"
)

type VirtualMemoryStat struct {
	PagesFree   int
	PagesActive int
}

type vmStatParser struct {
	buffer []byte
}

func (p *vmStatParser) Write(data []byte) (n int, err error) {
	p.buffer = append(p.buffer, data...)

	return len(data), nil
}

func (p *vmStatParser) Parse() (*VirtualMemoryStat, error) {
	s := bufio.NewScanner(bytes.NewBuffer(p.buffer))
	res := &VirtualMemoryStat{}

	for s.Scan() {
		key, value := parseLine(s.Text())

		switch key {
		case "Pages free":
			res.PagesFree = value
		case "Pages active":
			res.PagesActive = value
		}
	}

	return res, nil
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

func GetVirtualMemoryStat(ctx context.Context) *VirtualMemoryStat {
	parser := &vmStatParser{
		buffer: make([]byte, 1136),
	}

	cmd := exec.CommandContext(ctx, "vm_stat")
	cmd.Stdout = parser
	cmd.Run()

	res, _ := parser.Parse()

	return res
}
