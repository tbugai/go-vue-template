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

  http.HandleFunc("/", rootHandler)
  http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
  http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello from %s\n", r.URL.Path)
    log.Println(r.RemoteAddr, r.Method, r.URL)
  })

	err := http.ListenAndServe(":"+strconv.FormatInt(s.Port, 10), nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func rootHandler(rw http.ResponseWriter, r *http.Request) {
  http.ServeFile(rw, r, "backend/templates/index.html")
}
