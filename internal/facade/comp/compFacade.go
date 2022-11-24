package comp

import (
	"context"
	"github.com/calebtracey/rugby-crawler-api/external/models/request"
	"github.com/calebtracey/rugby-crawler-api/external/models/response"
	"github.com/calebtracey/rugby-crawler-api/internal/dao/comp"
	compCrawl "github.com/calebtracey/rugby-crawler-api/internal/dao/comp"
	"github.com/calebtracey/rugby-crawler-api/internal/dao/psql"
	log "github.com/sirupsen/logrus"
)

const LeagueIDSixNations = "180659"

//go:generate mockgen -destination=mockFacade.go -package=comp . FacadeI
type FacadeI interface {
	CompetitionCrawl(ctx context.Context, req request.CompetitionCrawlRequest) (resp response.CompetitionCrawlResponse)
}

type Facade struct {
	DBDAO      psql.DAOI
	CompDAO    compCrawl.DAOI
	CompMapper comp.MapperI
}

func (s Facade) CompetitionCrawl(ctx context.Context, req request.CompetitionCrawlRequest) (resp response.CompetitionCrawlResponse) {
	//TODO create scrape url
	url := ""
	resp, err := s.CompDAO.CompCrawlData(ctx, url, req.Date)
	if err != nil {

		log.Error(err)
		return response.CompetitionCrawlResponse{
			Message: response.Message{
				ErrorLog: response.ErrorLogs{
					*err,
				},
			},
		}
	}
	compExec := s.CompMapper.MapAddCompetitionData(resp.CompId, resp.Name, resp.TeamIds)
	_, dbErr := s.DBDAO.InsertOne(ctx, compExec)
	if dbErr != nil {
		log.Error(dbErr)
		return response.CompetitionCrawlResponse{
			Message: response.Message{
				ErrorLog: response.ErrorLogs{
					*dbErr,
				},
			},
		}
	}
	//TODO add response mapping
	return resp
}
