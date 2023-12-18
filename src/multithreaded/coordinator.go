package MultiThreaded

import (
	"fmt"
	"sync"
)


type KeyValue struct {
	Key   string
	Value int
}


// MultiThreadedMR represents a single-threaded MapReduce instance
type MultiThreadedMR struct {
	Tasks []string
	NumReducers int
	Pipes []chan  KeyValue

}


func NewMultiThreadedMR(tasks []string) *MultiThreadedMR {
	return &MultiThreadedMR{
		Tasks:        tasks,
		NumReducers:   4, // Set your default value for NumReducers here
		Pipes: make([]chan KeyValue, 4), // Set your default value for ReducerPipes here
	}
}
//Start Processing
func (mr *MultiThreadedMR) Process() {
	var mapperwg sync.WaitGroup
	var reducerwg sync.WaitGroup
	for _,e :=range(mr.Tasks) {

	}

	for i := 0 ;i < mr.NumReducers ; i++ {
		reducerwg.Add(1)
		go reduceFunction(mr.ReducerPipes[i])
	}




}




func main() {
	// Example usage
	mr := NewMultiThreadedMR{
		Tasks: []string{"file1.txt", "file2.txt", "file3.txt"},
	}
	result := mr.Process()
	fmt.Println("Aggregate Result:", result)
}



