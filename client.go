package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	go receiveMessages(conn)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		message := scanner.Text()
		if message == "/exit" {
			break
		}
		fmt.Fprintln(conn, message)
	}
}

func receiveMessages(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error receiving message:", err)
			break
		}
		fmt.Print(message)
	}
}
