package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/F7icK/check_sites_timings/internal/check_sites_timings/service"
	"github.com/F7icK/check_sites_timings/pkg/customerr"
	"github.com/F7icK/check_sites_timings/pkg/customlog"
)

type Handlers struct {
	s *service.Service
}

func NewHandlers(s *service.Service) *Handlers {
	return &Handlers{
		s: s,
	}
}

type result struct {
	Err string `json:"error"`
}

func apiErrorEncode(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if customError, ok := err.(*customerr.CustomError); ok {
		w.WriteHeader(customError.Code)
	}

	r := result{Err: err.Error()}

	if err = json.NewEncoder(w).Encode(r); err != nil {
		customlog.Error.Println(err)
	}
}

func apiResponseEncoder(w http.ResponseWriter, res interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		customlog.Error.Println(err)
	}
}
