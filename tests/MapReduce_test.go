package MapReduce

import (
	"testing"
	"MapReduce/MultiThreaded"
	"MapReduce/SingleThreaded"
	"time"
)



// 
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


func runSingleThreadedTest(files []string, t *testing.T) time.Duration {
	startTime := time.Now()
	singleThreadedMR := SingleThreaded.NewSingleThreadedMR(files)
	_ = singleThreadedMR.Process()
	duration := time.Since(startTime)
	t.Logf("Single-threaded processing took: %v", duration)
	return duration
}

// Function to run a multi-threaded test
func runMultiThreadedTest(files []string, numReducers int, t *testing.T) time.Duration {
	startTime := time.Now()
	mr := MultiThreaded.NewMultiThreadedMR(numReducers)
	_ = mr.Process(files)
	duration := time.Since(startTime)
	t.Logf("Multi-threaded processing took: %v", duration)
	return duration
}

// Test to check the time taken by single-threaded processing
func TestSingleThreadedMRProcessing(t *testing.T) {
	var totalSingleThreaded time.Duration
	// Run the test multiple times to find the average increase
	for i := 0; i < 5; i++ {
		files := []string{
			"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt"}
		totalSingleThreaded += runSingleThreadedTest(files, t)
	}

	averageSingleThreaded := totalSingleThreaded / 5
	t.Logf("Average time for single-threaded processing: %v", averageSingleThreaded)
	
}

// Test to check the time taken by multi-threaded processing
func TestMultiThreadedMRProcessing(t *testing.T) {
	var totalMultiThreaded time.Duration

	// Run the test multiple times to find the average increase
	for i := 0; i < 5; i++ {
		numReducers := 4
		files := []string{"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt"}
		totalMultiThreaded += runMultiThreadedTest(files, numReducers, t)
	}

	averageMultiThreaded := totalMultiThreaded / 5
	t.Logf("Average time for multi-threaded processing: %v", averageMultiThreaded)
}
