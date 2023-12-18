package MultiThreaded
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



	