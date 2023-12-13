package DummyJSONWebserver

import (
    "fmt"
    "net"
    "strings"
)



type GreetingResponse struct {
    Greeting string `json:"greeting"`
}

type Handler[T any] func() T

type ResponseWriter struct {
    conn net.Conn
}

func (w *ResponseWriter) Write(data []byte) (int, error) {
    return w.conn.Write(data)
}

type Router struct {
    routes map[string]interface{}
}

type RouteRegister struct {
    router *Router
}

func NewRouteRegister() *RouteRegister {
    return &RouteRegister{
        router: &Router{
            routes: make(map[string]interface{}),
        },
    }
}

func (rr *RouteRegister) RegisterRoute(route string, handler func() interface{}) {
    rr.router.routes[route] = handler
}

func (register *RouteRegister) startListening(port string) {
    listener, err := net.Listen("tcp", ":"+port)
    if err != nil {
        fmt.Println("Error listening:", err)
        return
    }
    defer listener.Close()
    fmt.Printf("Server listening on :%s\n", port)

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        go handleConnection(conn,register)
    }
}
func handleConnection(conn net.Conn, register *RouteRegister) {
    defer conn.Close()
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Error reading:", err)
        return
    }
    request := string(buffer[:n])
    path := extractPath(request)

    // Assume path is "/greeting" for this example
    if handler, ok := register.router.routes[path]; ok {
        // Execute the handler function and get the response
        response := handler.(func() interface{})()
        // Convert the response to a JSON string
        responseJSON := fmt.Sprintf(`{"greeting": "%s"}`, response.(GreetingResponse).Greeting)
        // Construct the HTTP response
        httpResponse := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: application/json\r\nContent-Length: %d\r\n\r\n%s", len(responseJSON), responseJSON)
        // Send the response
        conn.Write([]byte(httpResponse))
    }
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

func GreetingHandler() interface{} {
    return GreetingResponse{Greeting: "Hello, world!"}
}



