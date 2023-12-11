package loadbalancer

import (
	"sync"
)

// Service holder, Services are responsible for registering themselves, This is done on server startup.

type ServiceRegistry struct {
	mu sync.Mutex
	servers map[Server]struct{}
}


type Server struct {
	ip string
	port string
}

func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{
		servers: make(map[Server]struct{}),
	}	
}
func (registry *ServiceRegistry) RegisterServer(server Server) {
	registry.mu.Lock()
	defer registry.mu.Unlock()
	registry.servers[server] = struct{}{}
} 

func (registry *ServiceRegistry) DeRegisterServer(server Server) {
	registry.mu.Lock()
	defer registry.mu.Unlock()
	delete(registry.servers, server)
} 









