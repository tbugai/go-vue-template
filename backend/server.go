package backend

import (
    "fmt"
)

type Server struct {
    Port int32
}

func (s Server) Serve() {
    fmt.Print("Serving on port: ", s.Port)
}
