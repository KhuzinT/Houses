package rest

import (
	"Houses/internal/service"
	"Houses/internal/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	s *service.Service
	a *utils.AuthManager
}

func NewHandler(s *service.Service, a *utils.AuthManager) *Handler {
	return &Handler{s: s, a: a}
}

func (h *Handler) NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/login", h.login).Methods("POST")
	r.HandleFunc("/register", h.register).Methods("POST")

	r.HandleFunc("/house/{id}", h.getFlats).Methods("GET")
	r.HandleFunc("/house/create", h.createHouse).Methods("POST")

	r.HandleFunc("/flat/create", h.createFlat).Methods("POST")
	r.HandleFunc("/flat/update", h.updateFlat).Methods("POST")

	return r
}
