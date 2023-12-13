package DummyJSONWebserver

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type LoadBalancer struct {
	serverStatus map[string]bool
	servers      []string
	current      int
	mutex        sync.Mutex
	client       *http.Client
}

func NewLoadBalancer(listofservers []string) *LoadBalancer {
	// Initialize serverStatus map with all servers marked as healthy
	serverStatus := make(map[string]bool)
	for _, server := range listofservers {
		serverStatus[server] = true
	}

	return &LoadBalancer{
		serverStatus: serverStatus,
		servers:      listofservers,
		current:      0,
		client:       &http.Client{Timeout: 1 * time.Second},
	}
}

func (lb *LoadBalancer) markServerAsDown(server string) {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()
	lb.serverStatus[server] = false
}

func (lb *LoadBalancer) markServerAsUp(server string) {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()
	lb.serverStatus[server] = true
}

func (lb *LoadBalancer) getNextBackend() string {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()
	// Iterate through servers to find the next healthy server
	for i := 0; i < len(lb.servers); i++ {
		backend := lb.servers[lb.current]
		lb.current = (lb.current + 1) % len(lb.servers)
		if lb.serverStatus[backend] {
			return backend
		}
	}
	return "" // No healthy backend found
}

func (lb *LoadBalancer) handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request from %s\n", r.RemoteAddr)

	backend := lb.getNextBackend()

	if backend == "" {
		http.Error(w, "All backend servers are unhealthy", http.StatusServiceUnavailable)
		return
	}

	// Forward the request to the selected healthy backend server
	resp, err := http.Get(backend)
	if err != nil {
		// If the request to the backend fails, mark the server as down
		lb.markServerAsDown(backend)
		http.Error(w, "Error forwarding request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// If the backend responds successfully, mark the server as up
	lb.markServerAsUp(backend)

	// Copy the backend server's response to the client
	io.Copy(w, resp.Body)
}
