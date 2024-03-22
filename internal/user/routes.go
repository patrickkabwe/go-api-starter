package user

import (
	"fmt"
	"github.com/patrickkabwe/go-api-starter/utils"
	"net/http"
)

type Handler struct {
	store Store
}

func NewHandler(store Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r *http.ServeMux) {
	p := "/api/v1"
	r.HandleFunc(fmt.Sprintf("%s/users/me", p), h.handleGetUser)
}

func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"id": "test",
	})
	return
}
