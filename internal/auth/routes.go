package auth

import (
	"fmt"
	"github.com/patrickkabwe/go-api-starter/internal/user"
	"github.com/patrickkabwe/go-api-starter/types"
	"github.com/patrickkabwe/go-api-starter/utils"
	"net/http"
)

type Handler struct {
	store user.Store
}

func NewHandler(store user.Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r *http.ServeMux) {
	p := "/api/v1"
	r.HandleFunc(fmt.Sprintf("POST %s/register", p), h.handleRegister)
	r.HandleFunc(fmt.Sprintf("POST %s/login", p), h.handleLogin)
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var registerPayload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &registerPayload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Create user
	err := h.store.CreateUser(&types.User{
		FirstName: registerPayload.FirstName,
		LastName:  registerPayload.LastName,
		Email:     registerPayload.Email,
		Password:  registerPayload.Password,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"token": "token"})
	return
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

	var loginPayload types.LoginUserPayload
	if err := utils.ParseJSON(r, &loginPayload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": "token"})
	return
}
