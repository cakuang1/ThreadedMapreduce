package main

import (
	"fmt"
	"runtime"
)

func main() {
	numCores := runtime.NumCPU()
	fmt.Printf("Number of CPU cores: %d\n", numCores)
}
