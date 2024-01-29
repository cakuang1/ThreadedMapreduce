package main

import (
	"fmt"
	"./MultiThreaded" // Update this import path based on your project structure
)

func main() {
	// Example usage
	numReducers := 4
	mr := MultiThreaded.NewMultiThreadedMR(numReducers)

	files := []string{"file1.txt", "file2.txt", "file3.txt"}
	mr.Process(files)

	fmt.Println("Processing completed.")
}
