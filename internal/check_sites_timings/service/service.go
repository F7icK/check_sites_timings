package service

import (
	"net/http"
	"time"

	"github.com/F7icK/check_sites_timings/internal/check_sites_timings/service/checker"
	"github.com/F7icK/check_sites_timings/internal/check_sites_timings/service/statistics"
	"github.com/F7icK/check_sites_timings/internal/check_sites_timings/types"
	"github.com/F7icK/check_sites_timings/internal/clients/repository"
	"github.com/F7icK/check_sites_timings/pkg/customlog"
)

type Checker interface {
	CheckAllSites(httpClient *http.Client) error
	CheckSite(newSite string) (*types.Site, error)
	GetMinRequestTime() (*types.Site, error)
	GetMaxRequestTime() (*types.Site, error)
	AddHistory(endpoint string) (*types.History, error)
}

type Statistics interface {
	GetStatistics() ([]types.Statistics, error)
}

type Service struct {
	Checker
	Statistics
}

func NewService(db *repository.Repository, chanStop <-chan struct{}) *Service {
	cc := &Service{
		Checker:    checker.NewCheckerService(db.Checker),
		Statistics: statistics.NewStatisticsService(db.Statistics),
	}

	go customWorker(cc, chanStop)

	return cc
}

func customWorker(service *Service, chanStop <-chan struct{}) {
	httpClient := http.Client{
		Timeout: 50 * time.Second,
	}

	if err := service.CheckAllSites(&httpClient); err != nil {
		customlog.Error.Println(err)
	}

	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := service.CheckAllSites(&httpClient); err != nil {
				customlog.Error.Println(err)
				return
			}
		case <-chanStop:
			customlog.Info.Println("customWorker stopped")
			return
		}
	}
}
