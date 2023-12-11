package webserver



import (
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	request := string(buffer[:n])
	path := extractPath(request)
	response := fmt.Sprintf("HTTP/1.1 200 OK\r\n\r\nRequested path: %s\r\n", path)
	conn.Write([]byte(response))
}
func extractPath(request string) string {
	lines := strings.Split(request, "\n")
	if len(lines) > 0 {
		parts := strings.Fields(lines[0])
		if len(parts) > 1 {
			return parts[1]
		}
	}
	return "/"
}


func main() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on :80")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}
