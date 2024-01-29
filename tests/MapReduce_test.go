package MapReduce

import (
	"testing"
	"MapReduce/MultiThreaded"
	"MapReduce/SingleThreaded"
)


// Test to check if the single-threaded and multi-threaded versions produce the same outcome
func TestMRProcessing(t *testing.T) {
	// Single-threaded processing
	singleThreadedMR := SingleThreaded.NewSingleThreadedMR([]string{"a-hell0.txt","alien3a.txt","aliens3.txt", "aliensfaq.txt"})
	singleThreadedResult := singleThreadedMR.Process()

	// Multi-threaded processing
	numReducers := 4
	mr := MultiThreaded.NewMultiThreadedMR(numReducers)
	files := []string{"a-hell0.txt","alien3a.txt","aliens3.txt", "aliensfaq.txt"}
	multiThreadedResult := mr.Process(files)
	// Compare the results
	if !compareResults(singleThreadedResult, multiThreadedResult) {
		t.Errorf("Results do not match. Single-threaded result: %v, Multi-threaded result: %v", singleThreadedResult, multiThreadedResult)
	}
}
// Helper function to compare two results

func compareResults(result1, result2 map[string]int) bool {
	// Compare the length of the maps
	if len(result1) != len(result2) {
		return false
	}

	// Compare the content of the maps
	for key, value := range result1 {
		if result2Value, ok := result2[key]; !ok || value != result2Value {
			return false
		}
	}

	return true
}
