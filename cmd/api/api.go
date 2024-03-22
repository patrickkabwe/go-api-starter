package api

import (
	"github.com/patrickkabwe/go-api-starter/internal/auth"
	"github.com/patrickkabwe/go-api-starter/internal/user"
	"github.com/patrickkabwe/go-api-starter/utils"
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

	mux.HandleFunc("/api/v1/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		utils.WriteJSON(w, http.StatusOK, map[string]string{"status": "up and running..."})
	})

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(mux)

	authHandler := auth.NewHandler(userStore)
	authHandler.RegisterRoutes(mux)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, mux)
}

func (s *APIServer) Close() {

}
