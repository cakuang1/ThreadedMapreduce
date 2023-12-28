package tests

import (
	"fmt"
	"testing"
)


// Define your SingleThreadedMR and NewMultiThreadedMR types and their methods here.

// Example implementation of SingleThreadedMR and NewMultiThreadedMR
type SingleThreadedMR struct {
	Tasks []string
}



func (mr SingleThreadedMR) Process() string {
	// Implement the single-threaded processing logic
	return "SingleThreadedResult"
}

type MultiThreadedMR struct {
	Tasks []string
}

func NewMultiThreadedMR(tasks []string) *MultiThreadedMR {
	return &MultiThreadedMR{Tasks: tasks}
}

func (mr *MultiThreadedMR) Process() string {
	// Implement the multi-threaded processing logic
	return "MultiThreadedResult"
}

// Test to check if the single-threaded and multi-threaded versions produce the same outcome
func TestMRProcessing(t *testing.T) {
	// Single-threaded processing
	singleThreadedMR := SingleThreadedMR{Tasks: []string{"file1.txt", "file2.txt", "file3.txt"}}
	singleThreadedResult := singleThreadedMR.Process()

	// Multi-threaded processing
	multiThreadedMR := NewMultiThreadedMR([]string{"file1.txt", "file2.txt", "file3.txt"})
	multiThreadedResult := multiThreadedMR.Process()

	// Compare the results
	if singleThreadedResult != multiThreadedResult {
		t.Errorf("Results do not match. Single-threaded result: %s, Multi-threaded result: %s", singleThreadedResult, multiThreadedResult)
	}
}


