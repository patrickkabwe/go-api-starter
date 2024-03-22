package api

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type APIServer struct {
	db   *gorm.DB
	addr string
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, mux)
}

func (s *APIServer) Close() {

}
