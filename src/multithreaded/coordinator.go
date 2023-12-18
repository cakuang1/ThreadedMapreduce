package MultiThreaded

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"log"
)


type KeyValue struct {
	Key   string
	Value int
}


// MultiThreadedMR represents a single-threaded MapReduce instance
type MultiThreadedMR struct {
	Tasks []string
	NumReducers int
	ReducerPipes []chan <- KeyValue
}


func NewMultiThreadedMR(tasks []string) *MultiThreadedMR {
	return &MultiThreadedMR{
		Tasks:        tasks,
		NumReducers:   4, // Set your default value for NumReducers here
		ReducerPipes: make([]chan<- KeyValue, 4), // Set your default value for ReducerPipes here
	}
}






//Start Processing
func (mr *MultiThreadedMR) Process() {


}




func main() {
	// Example usage
	mr := NewMultiThreadedMR{
		Tasks: []string{"file1.txt", "file2.txt", "file3.txt"},
	}
	// In a multithreaded env, the number of mapper threads will be equal to the number of files
	// Reducers ar
	// Process all files and aggregate the results
	result := mr.Process()
	fmt.Println("Aggregate Result:", result)
}



