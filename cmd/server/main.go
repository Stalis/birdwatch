package main

import (
	"fmt"

	"github.com/Stalis/birdwatch/pkg/stat"
)

func main() {
	info, err := stat.GetInfo()
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

	fmt.Printf("CPU: %v\n", info.CPULoad)
	fmt.Printf("Memory: %v\n", info.Memory)
}
