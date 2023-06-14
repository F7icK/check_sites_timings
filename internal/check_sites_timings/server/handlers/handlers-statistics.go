package handlers

import "net/http"

func (h *Handlers) Statistics(w http.ResponseWriter, _ *http.Request) {
	statistics, err := h.s.GetStatistics()
	if err != nil {
		apiErrorEncode(w, err)
		return
	}

	apiResponseEncoder(w, statistics)
}
