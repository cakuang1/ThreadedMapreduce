package main

import (
	"fmt"
	"MapReduce/MultiThreaded"

)

func main() {
	// Example usage
	numReducers := 4
	mr := MultiThreaded.NewMultiThreadedMR(numReducers)

	files := []string{"a-hell0.txt","alien3a.txt","aliens3.txt", "aliensfaq.txt"}
	mr.Process(files)

	fmt.Println("Processing completed.")
}
