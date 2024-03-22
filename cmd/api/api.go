package api

import (
	"fmt"
	"net/http"
)

type APIServer struct {
}

func NewAPIServer() *APIServer {
	return &APIServer{}
}

func (s *APIServer) Start(port string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	return http.ListenAndServe(port, mux)
}

func (s *APIServer) Close() {

}
