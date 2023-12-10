package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
	"sync"
)


// Represents a client who is TCP connected
type Client struct {
	conn   net.Conn
	name   string
	writer *bufio.Writer
}


type Message struct {
	Sender string
	Content string
	Time time.Time
}




type MessageQueue struct {
	messages []Message
	capacity int
	lock sync.Mutex

}


var clients = make(map[net.Conn]Client)




func main() {
	messagebuffer := make(chan Message, 20)

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Chat server is running on port 8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}




func handleConnection(conn net.Conn) {
	defer conn.Close()
	client := Client{
		conn:   conn,
		writer: bufio.NewWriter(conn),
	}



	name, err := client.receiveMessage()
	if err != nil {
		fmt.Println("Error receiving name:", err)
		return
	}
	client.name = strings.TrimSpace(name)
	clients[conn] = client
	broadcastMessage(client.name + " has joined the chat")
	// Handle incoming messages from the client
	for {
		message, err := client.receiveMessage()
		if err != nil {
			fmt.Println("Error receiving message:", err)
			break
		}

		if strings.TrimSpace(message) == "/exit" {
			delete(clients, conn)
			broadcastMessage(client.name + " has left the chat")
			break
		}

		broadcastMessage(client.name + ": " + message)
	}
}




func broadcastMessage(message string) {
	fmt.Println(message)
	for _, client := range clients {
		client.writer.WriteString(message + "\n")
		client.writer.Flush()
	}
}






func (c *Client) sendMessage(message string) {
	c.writer.WriteString(message)
	c.writer.Flush()
}



func (c *Client) receiveMessage() (string, error) {
	message, err := bufio.NewReader(c.conn).ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(message), nil
}




