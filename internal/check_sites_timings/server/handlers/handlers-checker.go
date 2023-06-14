package handlers

import (
	"net/http"
)

func (h *Handlers) CheckSite(w http.ResponseWriter, r *http.Request) {
	site := r.FormValue("site")

	info, err := h.s.CheckSite(site)
	if err != nil {
		apiErrorEncode(w, err)
		return
	}

	apiResponseEncoder(w, info)
}

func (h *Handlers) CheckSiteMinTime(w http.ResponseWriter, r *http.Request) {
	site, err := h.s.GetMinRequestTime()
	if err != nil {
		apiErrorEncode(w, err)
		return
	}

	apiResponseEncoder(w, site)
}

func (h *Handlers) CheckSiteMaxTime(w http.ResponseWriter, r *http.Request) {
	site, err := h.s.GetMaxRequestTime()
	if err != nil {
		apiErrorEncode(w, err)
		return
	}

	apiResponseEncoder(w, site)
}
