package MultiThreaded

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"os"
	"strings"
	"sync"
)

type KeyValue struct {
	Key   string
	Value int
}

type MultiThreadedMR struct {
	NumReducers int
	Pipes       []chan KeyValue
	ReducerToOutput chan map[string]int
	MapperWG    sync.WaitGroup
	ReducerWG   sync.WaitGroup
}

// 
func NewMultiThreadedMR(numReducers int) *MultiThreadedMR {
	defaultBufferSize := 10
	mr := &MultiThreadedMR{
		NumReducers:      numReducers,
		Pipes:            make([]chan KeyValue, numReducers),
		ReducerToOutput:  make(chan map[string]int, numReducers), 
	}
	for i := 0; i < numReducers; i++ {
		mr.Pipes[i] = make(chan KeyValue, defaultBufferSize)
	}
	
	return mr
}



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


func (mr *MultiThreadedMR) Reduce(input chan KeyValue) map[string]int {
	wordCounts := make(map[string]int)
	// Process key-value pairs from the input channel
	for kv := range input {
		wordCounts[kv.Key] += kv.Value
	}
	return wordCounts
}


func (mr *MultiThreadedMR) Mapper(lines []string) {
    defer mr.MapperWG.Done()
    mr.MapperWG.Add(1)

    output := make(map[string]int)
    for _, line := range lines {
        words := strings.Fields(line)
        for _, word := range words {
            output[word]++
        }
    }

    for key, word := range output {
        hash := hashString(key) % uint32(mr.NumReducers)
        mr.Pipes[hash] <- KeyValue{Key: key, Value: word}
    }
    fmt.Println("Mapper Done")
}

func (mr *MultiThreadedMR) Process(files []string) map[string]int {
    // Start mapper goroutines
    for _, fileName := range files {
        lines, err := ReadFileLineByLine(fileName)
        if err != nil {
            fmt.Println("Error reading file:", err)
            return nil
        }
        go mr.Mapper(lines)
    }

    // Close mapper channels to signal reducers to finish

    // Launch reducers
    for i := 0; i < mr.NumReducers; i++ {
        mr.ReducerWG.Add(1)
        go func(ch chan KeyValue, outputChan chan map[string]int) {
            defer mr.ReducerWG.Done()
            // Perform reduction and get the result
            result := mr.Reduce(ch)
            // Send the result to the output channel
            outputChan <- result
        }(mr.Pipes[i], mr.ReducerToOutput)
    }
    // Wait for all mapper and reducer goroutines to finish
    mr.MapperWG.Wait()

    // Close mapper channels to signal reducers to finish
    for i := 0; i < mr.NumReducers; i++ {
        close(mr.Pipes[i])
    }

    // Wait for all reducer goroutines to finish
    mr.ReducerWG.Wait()

    // Close the ReducerToOutput channel after collecting all results
    close(mr.ReducerToOutput)
    // Collect and concatenate the final results
    finalResult := make(map[string]int)
    for result := range mr.ReducerToOutput {
        for key, value := range result {
            finalResult[key] += value
        }
    }

    // Print the final result

    return finalResult
}



func hashString(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

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
