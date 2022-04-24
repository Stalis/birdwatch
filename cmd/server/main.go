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

	fmt.Println(info)
}
