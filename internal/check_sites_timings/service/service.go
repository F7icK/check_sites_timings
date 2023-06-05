package service

import (
	"github.com/F7icK/check_sites_timings/internal/check_sites_timings/service/checker"
	"github.com/F7icK/check_sites_timings/internal/check_sites_timings/service/statistics"
)

type Checker interface {
}

type Statistics interface {
}

type Service struct {
	Checker
	Statistics
}

func NewService() *Service {

	return &Service{
		Checker:    checker.NewCheckerService(),
		Statistics: statistics.NewStatisticsService(),
	}
}
