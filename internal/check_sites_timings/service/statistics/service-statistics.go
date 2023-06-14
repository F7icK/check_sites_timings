package statistics

import (
	"github.com/F7icK/check_sites_timings/internal/check_sites_timings/types"
	"github.com/F7icK/check_sites_timings/internal/clients/repository"
	"github.com/F7icK/check_sites_timings/pkg/customerr"
	"github.com/F7icK/check_sites_timings/pkg/customlog"
)

type StatService struct {
	db repository.Statistics
}

func NewStatisticsService(db repository.Statistics) *StatService {
	return &StatService{
		db: db,
	}
}

func (s *StatService) GetStatistics() ([]types.Statistics, error) {
	statistics, err := s.db.CountUseEndpoint()
	if err != nil {
		customlog.Error.Println(err)
		return nil, customerr.ErrorInternalServerError
	}

	return statistics, nil
}
