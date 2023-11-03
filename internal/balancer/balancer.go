package balancer

import (
	"fmt"
	"net/http"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(w http.ResponseWriter, r *http.Request)
}

type LoadBalancer struct {
	port    string
	counter int
	servers []Server
}

func New(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:    port,
		counter: 0,
		servers: servers,
	}
}

func (lb *LoadBalancer) Next() Server {
	server := lb.servers[lb.counter%len(lb.servers)]
	for !server.IsAlive() {
		lb.counter++
		server = lb.servers[lb.counter%len(lb.servers)]
	}
	lb.counter++
	return server
}

func (lb *LoadBalancer) ServeProxy(w http.ResponseWriter, r *http.Request) {
	targetServer := lb.Next()
	fmt.Printf("forwarding request to address %q \n", targetServer.Address())
	targetServer.Serve(w, r)
}
