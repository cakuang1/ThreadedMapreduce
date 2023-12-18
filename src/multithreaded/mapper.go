package MapReduce
// MAPPER FUNCTION Also includes a combiner

import (
	"fmt"
	"hash/fnv"
	"runtime"
	"strings"
	"sync"
)




func mapper(input string, output chan<- map[string]int) {
	wordCount := make(map[string]int)

	words := strings.Fields(input)
	for _, word := range words {
		wordCount[word]++
	}

	// Send the word count directly to the reducer
	output <- wordCount
}



	