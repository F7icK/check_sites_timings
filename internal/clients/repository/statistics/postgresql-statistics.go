package statistics

import (
	"github.com/F7icK/check_sites_timings/internal/check_sites_timings/types"
	"gorm.io/gorm"
)

type StatPostgres struct {
	db *gorm.DB
}

func NewStatPostgres(db *gorm.DB) *StatPostgres {
	return &StatPostgres{db: db}
}

func (p *StatPostgres) CountUseEndpoint() ([]types.Statistics, error) {
	statistics := make([]types.Statistics, 0)

	if err := p.db.Debug().Select("endpoint, count(id) AS count_use").Table("history").Group("endpoint").Find(&statistics).Error; err != nil {
		return nil, err
	}

	return statistics, nil
}
