package repository

import (
	"github.com/F7icK/check_sites_timings/internal/check_sites_timings/types"
	"github.com/F7icK/check_sites_timings/internal/clients/repository/checker"
	"github.com/F7icK/check_sites_timings/internal/clients/repository/statistics"
	"gorm.io/gorm"
)

type Checker interface {
	Begin() *gorm.DB
	GetAllSites() ([]types.Site, error)
	UpdateSite(site types.Site) (*types.Site, error)
	GetSiteByName(nameSite string) (*types.Site, error)
	AddSite(site *types.Site) (*types.Site, error)
	GetMinReqTime() (*types.Site, error)
	GetMaxReqTime() (*types.Site, error)
	AddHistory(data *types.History) (*types.History, error)
}

type Statistics interface {
	CountUseEndpoint() ([]types.Statistics, error)
}

type Repository struct {
	Checker
	Statistics
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Checker:    checker.NewCheckPostgres(db),
		Statistics: statistics.NewStatPostgres(db),
	}
}
