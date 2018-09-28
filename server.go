package main

import (
    "template/backend"
)

func main() {
    server := backend.Server{Port: 3000}
    server.Serve()
}
