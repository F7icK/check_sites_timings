package handlers

import "github.com/F7icK/check_sites_timings/internal/check_sites_timings/service"

type Handlers struct {
	s *service.Service
}

func NewHandlers(s *service.Service) *Handlers {
	return &Handlers{
		s: s,
	}
}
