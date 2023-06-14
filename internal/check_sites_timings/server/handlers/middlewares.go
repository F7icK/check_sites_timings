package handlers

import (
	"net/http"
)

func (h *Handlers) AddHistory(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := h.s.AddHistory(r.URL.Path); err != nil {
			apiErrorEncode(w, err)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
