package MultiThreaded



func (mr *MultiThreadedMR) reduceFunction(input chan KeyValue) map[string]int {
	wordCounts := make(map[string]int)
	// Process key-value pairs from the input channel
	for kv := range input {
		wordCounts[kv.Key] += kv.Value
	}
	return wordCounts
}






