package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/F7icK/check_sites_timings/internal/check_sites_timings/server"
	"github.com/F7icK/check_sites_timings/internal/check_sites_timings/server/handlers"
	"github.com/F7icK/check_sites_timings/internal/check_sites_timings/service"
)

func main() {
	services := service.NewService()

	endpoints := handlers.NewHandlers(services)

	srv := server.NewServer(endpoints)

	stopFunc := func() {
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("err with Shutdown in main :%s", err.Error())
		}
	}

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	defer close(signalCh)

	go func(signalCh <-chan os.Signal, stopFunc func()) {
		select {
		case sig := <-signalCh:
			log.Printf("stopped with signal: %s", sig)
			stopFunc()
			os.Exit(0)
		}
	}(signalCh, stopFunc)

	if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
		log.Printf("err with Run in main :%s", err.Error())
	}
}
