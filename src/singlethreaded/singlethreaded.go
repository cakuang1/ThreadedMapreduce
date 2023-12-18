package MapReduce

// Single threaded map reduce



import (
	"fmt"
	"strings"
)

type KeyValue struct {
	Key   string
	Value int
}





func mapFunction(input string) []KeyValue {
	// Split the input into words
	words := strings.Fields(input)

	// Emit key-value pairs for each word
	var output []KeyValue
	for _, word := range words {
		output = append(output, KeyValue{Key: word, Value: 1})
	}

	return output
}

func reduceFunction(key string, values []int) KeyValue {
	count := 0
	for _, v := range values {
		count += v
	}

	return KeyValue{Key: key, Value: count}
}

func main() {
	// Input data
	inputData := "Hello world hello"

	// Map Phase
	mapOutput := mapFunction(inputData)

	// Shuffle and Sort Phase (omitted for simplicity in single-threaded version)

	// Reduce Phase
	result := make(map[string][]int)
	for _, kv := range mapOutput {
		result[kv.Key] = append(result[kv.Key], kv.Value)
	}

	// Final Output
	for key, values := range result {
		reduceOutput := reduceFunction(key, values)
		fmt.Printf("%s: %d\n", reduceOutput.Key, reduceOutput.Value)
	}
}
