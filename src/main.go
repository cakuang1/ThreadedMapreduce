package main



import (
	"fmt"
	"os"
)

func main() {
	// Check if there is exactly one command-line argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: program_name [option]")
		return

	}
	// Extract the command-line argument
	option := os.Args[1]
	// Check the value of the argument
	if option == "test1" {
		fmt.Println("You selected option 1.")
	} else if option == "test2" {
		fmt.Println("You selected option 2.")
	} else {
		fmt.Println("Invalid option. Please choose either option1 or option2.")
	}
}
