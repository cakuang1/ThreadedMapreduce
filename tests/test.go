package tests




import (
	"testing"
)



/* One simple test case that determines if the key value pairs returned from single threaded version
is the same as the multithreaded version*/





func main() {
	// Example usage
	mr := SingleThreadedMR{
		Tasks: []string{"file1.txt", "file2.txt", "file3.txt"},
	}
	// Process all files and aggregate the results
	result := mr.Process()
	fmt.Println("Aggregate Result:", result)
}



func main() {
	// Example usage
	mr := NewMultiThreadedMR([]string{"file1.txt", "file2.txt", "file3.txt"})
	mr.Process()
	fmt.Println("Processing completed.")
}
