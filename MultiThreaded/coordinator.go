package MultiThreaded

import (
	"fmt"
	"sync"
)

type KeyValue struct {
	Key   string
	Value int
}

type MultiThreadedMR struct {
	Tasks []string
	NumReducers int
	Pipes []chan  KeyValue
	MapperWG   sync.WaitGroup
	ReducerWG  sync.WaitGroup
}



func NewMultiThreadedMR(tasks []string) *MultiThreadedMR {
	numReducers := 4
	defaultBufferSize := 10 
	mr := &MultiThreadedMR{
		Tasks:        tasks,
		NumReducers:  numReducers,
		Pipes:        make([]chan KeyValue, numReducers),
	}
	for i := 0; i < numReducers; i++ {
		mr.Pipes[i] = make(chan KeyValue, defaultBufferSize)
	}
	return mr
}


func (mr *MultiThreadedMR) Process() {
	for _, e := range mr.Tasks {
		mr.MapperWG.Add(1)
		go mr.mapper(e)
	}
	for i := 0; i < mr.NumReducers; i++ {
		mr.ReducerWG.Add(1)
		go mr.reduceFunction(mr.Pipes[i])
	}
	// Wait for all mapper and reducer goroutines to finish
	mr.MapperWG.Wait()
	for i := 0; i < mr.NumReducers; i++ {
		close(mr.Pipes[i]) // Close the channel to signal reducers to finish
	}
	mr.ReducerWG.Wait()
}




func main() {
	// Example usage
	mr := NewMultiThreadedMR([]string{"file1.txt", "file2.txt", "file3.txt"})
	mr.Process()
	fmt.Println("Processing completed.")
}


