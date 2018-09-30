package backend

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	Port int64
}

func (s Server) Serve() {
  fmt.Printf("Listening on port: %d\n", s.Port)

  handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello from %s\n", r.URL.Path)
    log.Println(r.RemoteAddr, r.Method, r.URL)
  })

	err := http.ListenAndServe(":"+strconv.FormatInt(s.Port, 10), handler)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
