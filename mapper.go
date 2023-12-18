package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
)

const (
	numMappers  = 3
	numReducers = 2
)

func main() {
	// Read input file
	filePath := "input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	// Initialize channels for communication
	mapChannel := make(chan map[string]int, numMappers)
	reduceChannel := make(chan map[string]int, numReducers)

	// Start mapper nodes
	var wg sync.WaitGroup
	for i := 0; i < numMappers; i++ {
		wg.Add(1)
		go runMapper(i, file, mapChannel, &wg)
	}
	// Close mapChannel after all mappers are done
	go func() {
		wg.Wait()
		close(mapChannel)
	}()

	// Start reducer nodes
	for i := 0; i < numReducers; i++ {
		go runReducer(i, mapChannel, reduceChannel)
	}

	// Close reduceChannel after all reducers are done
	go func() {
		wg.Wait()
		close(reduceChannel)
	}()

	// Collect results from reducers
	finalResult := make(map[string]int)
	for result := range reduceChannel {
		mergeMaps(finalResult, result)
	}

	// Display final word count
	fmt.Println("Final Word Count:")
	for word, count := range finalResult {