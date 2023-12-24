# Go Multi-Threaded Single Machine MapReduce Framework


## Overview
This project provides a basic implementation of a multi-threaded MapReduce framework in a single Golang process using Golang Threads (goroutines). The goal of this project is to simulate a typical clustered algorithm on a single machine, similating each machine using a thread, and networking (mappers to reducers) through Go Channels. 






## Features
- **Multi-Threaded Processing:** Utilizes goroutines in Go for concurrent execution of mapping and reducing tasks.
- **Configurable:** Allows users to specify the number of reducers and customize the buffer size for communication channels.
- **Flexible Implementation:** Users can define their own mapping and reducing logic based on the specific requirements of their MapReduce application.





## Usage
1. **Instantiate MultiThreadedMR:**
    ```go
    mr := NewMultiThreadedMR([]string{"file1.txt", "file2.txt", "file3.txt"})
    ```

2. **Process Data:**
    ```go
    mr.Process()
    ```

3. **Implement Custom Logic:**
   - Fill in the `mapper` and `reduceFunction` methods with application-specific mapping and reducing logic.

4. **Results:**
   - The final aggregated result can be obtained based on the user-defined logic in the `reduceFunction`.


## Example
```go
package main

import (
	"fmt"
	"github.com/your-username/MultiThreadedMapReduce"
)

func main() {
	// Example usage
	mr := MultiThreadedMapReduce.NewMultiThreadedMR([]string{"file1.txt", "file2.txt", "file3.txt"})
	mr.Process()
	fmt.Println("Processing completed.")
}
```


## How this works
Typical multi-threaded processing framework with master-worker pattern. Threads are s wer 








 Master thread distributes tasks and maintains status and overall process.





