package server

import (
	"context"
	"net/http"
	"time"

	"github.com/F7icK/check_sites_timings/internal/check_sites_timings/server/handlers"
	"github.com/F7icK/check_sites_timings/pkg/customlog"
	"github.com/rs/cors"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(handlers *handlers.Handlers) *Server {
	router := NewRouter(handlers)

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		AllowedHeaders: []string{
			"*",
		},
	})

	return &Server{
		httpServer: &http.Server{
			Addr:         ":8080",
			Handler:      corsOpts.Handler(router),
			ReadTimeout:  60 * time.Second,
			WriteTimeout: 60 * time.Second,
		},
	}
}

func (s *Server) Run() error {
	customlog.Info.Println("Server started!")

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
