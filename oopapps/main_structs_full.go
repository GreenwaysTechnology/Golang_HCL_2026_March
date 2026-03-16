package main

import (
	"fmt"
	"strconv"
)

type Server struct {
	host string
	port int
}

// constructor
func NewServer(host string, port int) *Server {
	return &Server{
		host: host,
		port: port}
}

// methods
func (s *Server) Start() {
	fmt.Println("Server Started on", s.host, ":"+strconv.Itoa(s.port))
}

func main() {
	server := NewServer("localhost", 8080)
	server.port = 9000
	server.Start()
}
