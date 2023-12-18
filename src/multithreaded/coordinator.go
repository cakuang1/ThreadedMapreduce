package MapReduce


import (
	"fmt"
	"strings"
	"sync"
)




type MultiThreadedMapReduce struct {
	TaskQueue []string
	ReducerChannel []chan<- map[string]int
}



func (Mapper *MapReduce) txtFiles(files []string) {

	
}

func NewMapReducer

























