package MultiThreaded
// MAPPER FUNCTION Also includes a combiner


import (
	"bufio"
	"hash/fnv"
	"os"
	"strings"
	"fmt"
)



func hashString(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
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



	