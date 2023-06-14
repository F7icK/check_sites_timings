package checker

import (
	"github.com/F7icK/check_sites_timings/internal/check_sites_timings/types"
	"gorm.io/gorm"
)

type CheckPostgres struct {
	db *gorm.DB
}

func NewCheckPostgres(db *gorm.DB) *CheckPostgres {
	return &CheckPostgres{db: db}
}

func (p *CheckPostgres) Begin() *gorm.DB {
	tx := p.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	return tx
}

func (p *CheckPostgres) GetAllSites() ([]types.Site, error) {
	sites := make([]types.Site, 0)

	if err := p.db.Debug().Table("sites").Find(&sites).Error; err != nil {
		return nil, err
	}

	return sites, nil
}

func (p *CheckPostgres) UpdateSite(site types.Site) (*types.Site, error) {
	if err := p.db.Debug().Table("sites").Where("id = ?", site.ID).Updates(site).Error; err != nil {
		return nil, err
	}

	return &site, nil
}

func (p *CheckPostgres) GetSiteByName(nameSite string) (*types.Site, error) {
	site := new(types.Site)

	if err := p.db.Debug().Table("sites").Where("name = ?", nameSite).Take(&site).Error; err != nil {
		return nil, err
	}

	return site, nil
}

func (p *CheckPostgres) AddSite(site *types.Site) (*types.Site, error) {

	if err := p.db.Debug().Table("sites").Create(site).Error; err != nil {
		return nil, err
	}

	return site, nil
}

func (p *CheckPostgres) GetMinReqTime() (*types.Site, error) {
	site := new(types.Site)

	if err := p.db.Debug().Table("sites").Where("status_code = 200").Order("request_time_ms").Take(site).Error; err != nil {
		return nil, err
	}

	return site, nil
}

func (p *CheckPostgres) GetMaxReqTime() (*types.Site, error) {
	site := new(types.Site)

	if err := p.db.Debug().Table("sites").Where("status_code = 200").Order("request_time_ms desc").Take(site).Error; err != nil {
		return nil, err
	}

	return site, nil
}

func (p *CheckPostgres) AddHistory(data *types.History) (*types.History, error) {

	if err := p.db.Debug().Table("history").Create(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
