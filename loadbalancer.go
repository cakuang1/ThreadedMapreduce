package DummyJSONWebserver

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)	

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request from %s\n", r.RemoteAddr)
	// Forward the request to a backend server
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		http.Error(w, "Error forwarding request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	// Copy the backend server's response to the client
	io.Copy(w, resp.Body)
}

type LoadBalancer struct {
	servers []string
	current int
	mutex  sync.Mutex
}

func NewLoadBalancer(listofservers []string) *LoadBalancer{
	return &LoadBalancer{
		servers: listofservers,
		current: 0,
	}
}

func (lb *LoadBalancer) 









func main() {

	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":80", nil)
}
