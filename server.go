package gochat

// CHAT SERVER
import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// Represents a TCP connection from client
type Client struct {
	conn net.Conn
	name string
	writer *bufio.Writer 
}

// Hashtable for current connections
var clients = make(map[net.Conn]Client)

// server entrypoint
func main() {
	listener ,err := net.Listen("tcp","localhost:8080")
	if err != nil {
		fmt.Println("Error starting server : ",err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Chat server is running on port 8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection Error", err)
			continue
		}			

	}

}


// Handle the connection, each connection gets a goroutine
func handleConnection(conn net.Conn) {
	defer conn.Close()

	client := Client{
		conn : conn,
		writer: bufio.NewWriter(conn),
	}

	client.sendMessage("Enter your name to enter the chatroom :  ")
	name, err := client.receiveMessage()
	if err != nil {
		fmt.Println("Error receiving name:", err)
		return
	}
	
	client.name = strings.TrimSpace(name)
	clients[conn] = client
	broadcastMessage(client.name + " has joined the chat")
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
// Mapping for 




// Map 