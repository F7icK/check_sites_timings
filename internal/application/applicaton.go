package application

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/F7icK/check_sites_timings/internal/check_sites_timings/server"
	"github.com/F7icK/check_sites_timings/internal/check_sites_timings/server/handlers"
	"github.com/F7icK/check_sites_timings/internal/check_sites_timings/service"
	"github.com/F7icK/check_sites_timings/internal/check_sites_timings/types/config"
	"github.com/F7icK/check_sites_timings/internal/clients/repository"
	"github.com/F7icK/check_sites_timings/pkg/customlog"
	"github.com/F7icK/check_sites_timings/pkg/database/postgres"
)

func NewApplication(cfg *config.Config) {
	postgresClient, err := postgres.NewPostgres(cfg.PostgresDsn)
	if err != nil {
		customlog.Error.Println(err)
		return
	}

	db, err := postgresClient.Database()
	if err != nil {
		customlog.Error.Println(err)
		return
	}

	repos := repository.NewRepository(db)

	chanStop := make(chan struct{})

	services := service.NewService(repos, chanStop)

	endpoints := handlers.NewHandlers(services)

	srv := server.NewServer(endpoints)

	stopFunc := func() {
		close(chanStop)

		if err = srv.Shutdown(context.Background()); err != nil {
			customlog.Error.Println(err)
		}

		if err = postgresClient.Close(); err != nil {
			customlog.Error.Println(err)
		}
	}

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	defer close(signalCh)

	go func(signalCh <-chan os.Signal, stopFunc func()) {
		select {
		case sig := <-signalCh:
			customlog.Info.Printf("stopped with signal: %s", sig)
			stopFunc()
			os.Exit(0)
		}
	}(signalCh, stopFunc)

	if err = srv.Run(); !errors.Is(err, http.ErrServerClosed) {
		customlog.Error.Println(err)
		return
	}

	return
}
