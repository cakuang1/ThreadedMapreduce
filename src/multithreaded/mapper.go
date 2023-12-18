package MapReduce
// MAPPER FUNCTION Also includes a combiner
func mapFunction(input string) {
		words := strings.Fields(input)
		// Emit key-value pairs for each word
		var output []KeyValue
		for _, word := range words {
			output = append(output, KeyValue{Key: word, Value: 1})
		}

		return output
	}




	