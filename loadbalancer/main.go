package loadbalancer

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":80", nil)
}


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





