package checker

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/F7icK/check_sites_timings/internal/check_sites_timings/types"
	"github.com/F7icK/check_sites_timings/internal/clients/repository"
	"github.com/F7icK/check_sites_timings/pkg/customerr"
	"github.com/F7icK/check_sites_timings/pkg/customlog"
	"gorm.io/gorm"
)

type CheckService struct {
	db repository.Checker
}

func NewCheckerService(db repository.Checker) *CheckService {
	return &CheckService{
		db: db,
	}
}

func (s *CheckService) CheckAllSites(httpClient *http.Client) error {
	sites, err := s.db.GetAllSites()
	if err != nil {
		customlog.Error.Println(err)
		return err
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(sites))

	chanSite := make(chan types.Site, len(sites))

	for i := range sites {
		go newRequest(sites[i], httpClient, chanSite, wg)
	}

	go func() {
		wg.Wait()
		close(chanSite)
	}()

	for site := range chanSite {
		if _, err = s.db.UpdateSite(site); err != nil {
			customlog.Error.Println(err)
			return err
		}
	}

	return nil
}

func newRequest(dataSite types.Site, httpClient *http.Client, outDataSite chan<- types.Site, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()

	resp, err := httpClient.Get(fmt.Sprintf("https://%s", dataSite.Name))
	endTime := time.Since(start).Milliseconds()
	if err != nil {
		dataSite.StatusCode = http.StatusGatewayTimeout
	} else {
		resp.Body.Close()
		dataSite.StatusCode = resp.StatusCode
	}

	dataSite.RequestTimeMs = endTime

	outDataSite <- dataSite
}

func (s *CheckService) CheckSite(newSite string) (*types.Site, error) {
	uri, err := url.ParseRequestURI("https://" + newSite)
	if err != nil {
		return nil, customerr.ErrorBadRequest
	}

	if _, err = net.LookupHost(uri.Host); err != nil {
		return nil, customerr.ErrorBadRequest
	}

	oldSite, err := s.db.GetSiteByName(newSite)
	if err != nil && err != gorm.ErrRecordNotFound {
		customlog.Error.Println(err)
		return nil, customerr.ErrorInternalServerError
	}

	if oldSite != nil {
		return oldSite, nil
	}

	dataSite := &types.Site{Name: newSite}

	httpClient := http.Client{
		Timeout: 50 * time.Second,
	}

	start := time.Now()

	resp, err := httpClient.Get(fmt.Sprintf("https://%s", dataSite.Name))
	endTime := time.Since(start).Milliseconds()
	if err != nil {
		dataSite.StatusCode = http.StatusGatewayTimeout
	} else {
		resp.Body.Close()
		dataSite.StatusCode = resp.StatusCode
	}

	dataSite.RequestTimeMs = endTime

	if _, err = s.db.AddSite(dataSite); err != nil {
		customlog.Error.Println(err)
		return nil, customerr.ErrorInternalServerError
	}

	return dataSite, nil
}

func (s *CheckService) GetMinRequestTime() (*types.Site, error) {
	dataSite, err := s.db.GetMinReqTime()
	if err != nil {
		customlog.Error.Println(err)
		return nil, customerr.ErrorInternalServerError
	}

	return dataSite, nil
}

func (s *CheckService) GetMaxRequestTime() (*types.Site, error) {
	dataSite, err := s.db.GetMaxReqTime()
	if err != nil {
		customlog.Error.Println(err)
		return nil, customerr.ErrorInternalServerError
	}

	return dataSite, nil
}

func (s *CheckService) AddHistory(endpoint string) (*types.History, error) {
	history, err := s.db.AddHistory(&types.History{Endpoint: endpoint})
	if err != nil {
		customlog.Error.Println(err)
		return nil, customerr.ErrorInternalServerError
	}

	return history, nil
}
