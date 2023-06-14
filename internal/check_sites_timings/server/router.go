package server

import (
	"net/http"

	"github.com/F7icK/check_sites_timings/internal/check_sites_timings/server/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(h *handlers.Handlers) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	checker := router.PathPrefix("").Subrouter()
	checker.Use(h.AddHistory)
	checker.Methods(http.MethodPost).Path("/check").HandlerFunc(h.CheckSite)
	checker.Methods(http.MethodGet).Path("/check-min").HandlerFunc(h.CheckSiteMinTime)
	checker.Methods(http.MethodGet).Path("/check-max").HandlerFunc(h.CheckSiteMaxTime)

	//TODO полноценную админку не делал, так как по условию не сказано
	admin := router.PathPrefix("/admin").Subrouter()
	admin.Methods(http.MethodGet).Path("/statistics").HandlerFunc(h.Statistics)

	return router
}
