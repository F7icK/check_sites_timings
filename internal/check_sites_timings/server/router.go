package server

import (
	"github.com/F7icK/check_sites_timings/internal/check_sites_timings/server/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(h *handlers.Handlers) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	return router
}
