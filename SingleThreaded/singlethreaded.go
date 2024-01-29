package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// SingleThreadedMR represents a single-threaded MapReduce instance
type SingleThreadedMR struct {
	Tasks []string
}

// NewSingleThreadedMR creates a new SingleThreadedMR instance with the provided list of files
func NewSingleThreadedMR(files []string) *SingleThreadedMR {
	return &SingleThreadedMR{
		Tasks: files,
	}
}

// ReadFileLineByLine reads a file line by line
func ReadFileLineByLine(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// mapFunction performs the map operation for a list of lines
func mapFunction(lines []string) map[string]int {
	// Emit key-value pairs for each word in each line
	output := make(map[string]int)
	for _, line := range lines {
		words := strings.Fields(line)
		for _, word := range words {
			output[word]++
		}
	}

	return output
}

// reduceFunction performs the reduce operation on a list of key-value pairs
func reduceFunction(input map[string]int) map[string]int {
	// The reduce function can be a no-op for this example
	return input
}

// Process reads all files, performs the map and reduce operations, and aggregates the results
func (mr *SingleThreadedMR) Process() map[string]int {
	// Aggregate results from all files
	aggregateResult := make(map[string]int)

	// Map phase for all files
	for _, fileName := range mr.Tasks {
		lines, err := ReadFileLineByLine(fileName)
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue
		}
		// Map phase
		mapOutput := mapFunction(lines)

		// Aggregate map results
		for key, value := range mapOutput {
			aggregateResult[key] += value
		}
		log.Println("Mapping File : ", fileName)
	}
	// Reduce phase
	reduceOutput := reduceFunction(aggregateResult)
	return reduceOutput
}

func main() {
	// Example usage
	files := []string{"a-hell0.txt","alien3a.txt","aliens3.txt", "aliensfaq.txt"}
	mr := NewSingleThreadedMR(files)

	// Process all files and aggregate the results
	result := mr.Process()
	fmt.Println("Aggregate Result:", result)
}
