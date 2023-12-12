package webserver

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
)
type Handler func(w ResponseWriter, r *Request)
type Router struct {
	routes map[string]Handler
}
type ResponseWriter struct {
	conn net.Conn
}


func (w *ResponseWriter) Write(data []byte) (int, error) {
	return w.conn.Write(data)
}

func (router *Router) 


