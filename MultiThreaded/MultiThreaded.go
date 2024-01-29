package MultiThreaded

import (
	"fmt"
	"hash/fnv"
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
	MapperWG    sync.WaitGroup
	ReducerWG   sync.WaitGroup
}

func NewMultiThreadedMR(numReducers int) *MultiThreadedMR {
	defaultBufferSize := 10
	mr := &MultiThreadedMR{
		NumReducers: numReducers,
		Pipes:       make([]chan KeyValue, numReducers),
	}
	for i := 0; i < numReducers; i++ {
		mr.Pipes[i] = make(chan KeyValue, defaultBufferSize)
	}
	return mr
}

func (mr *MultiThreadedMR) mapper(input string) {
	defer mr.MapperWG.Done()
	fmt.Println("Mapping ", input)
	wordCount := make(map[string]int)
	words := strings.Fields(input)
	for _, word := range words {
		wordCount[word]++
	}
	for word, count := range wordCount {
		hash := hashString(word) % uint32(mr.NumReducers)
		mr.Pipes[hash] <- KeyValue{Key: word, Value: count}
	}
}

func (mr *MultiThreadedMR) reduceFunction(input chan KeyValue) map[string]int {
	wordCounts := make(map[string]int)
	// Process key-value pairs from the input channel
	for kv := range input {
		wordCounts[kv.Key] += kv.Value
	}
	return wordCounts
}

func (mr *MultiThreadedMR) Process(files []string) {
	for _, fileName := range files {
		mr.MapperWG.Add(1)
		go func(fileName string) {
			defer mr.MapperWG.Done()
			lines, err := ReadFileLineByLine(fileName)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
			for _, line := range lines {
				mr.mapper(line)
			}
		}(fileName)
	}

	// Launch reducers
	for i := 0; i < mr.NumReducers; i++ {
		mr.ReducerWG.Add(1)
		go func(ch chan KeyValue) {
			defer mr.ReducerWG.Done()
			result := mr.reduceFunction(ch)
			// Process the reduced data (in this example, just printing the results)
			fmt.Println("Reduced Result:", result)
		}(mr.Pipes[i])
	}

	// Wait for all mapper and reducer goroutines to finish
	mr.MapperWG.Wait()

	// Close mapper channels to signal reducers to finish
	for i := 0; i < mr.NumReducers; i++ {
		close(mr.Pipes[i])
	}

	mr.ReducerWG.Wait()
}

func hashString(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
