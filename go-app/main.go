package main

import (
	"log"
	"net/http"
)

type Server struct{}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "hello world 1234"}`))
}

func main() {
	s := &Server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":7070", nil))
}
