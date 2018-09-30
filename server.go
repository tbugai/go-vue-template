package main

import (
  "fmt"
  "flag"
	"template/backend"
)

var Port int64

func init() {
	flag.Int64Var(&Port, "port", 9000, "port to listen on")
}

func main() {
  fmt.Printf("Starting server...\n")
	server := backend.Server{Port: Port}
	server.Serve()
}
