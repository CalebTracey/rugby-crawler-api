package comp

import (
	"context"
	"github.com/calebtracey/rugby-crawler-api/external/models/response"
	"github.com/gocolly/colly"
)

//go:generate mockgen -destination=mockDao.go -package=comp . DAOI
type DAOI interface {
	CompCrawlData(ctx context.Context, url, date string) (resp response.CompetitionCrawlResponse, log *response.ErrorLog)
}

type DAO struct {
	Collector *colly.Collector
}

func (s DAO) CompCrawlData(ctx context.Context, url, date string) (resp response.CompetitionCrawlResponse, log *response.ErrorLog) {
	//s.Collector.OnHTML()
	return resp, nil
}

func mapError(err error, query string) (errLog *response.ErrorLog) {
	errLog = &response.ErrorLog{
		Query: query,
	}
	if err != nil {
		errLog.RootCause = err.Error()
	}
	errLog.StatusCode = "500"
	return errLog
}
