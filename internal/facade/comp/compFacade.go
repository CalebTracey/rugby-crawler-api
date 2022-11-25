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

//go:generate mockgen -destination=mockFacade.go -package=comp . FacadeI
type FacadeI interface {
	CrawlLeaderboard(ctx context.Context, req request.CrawlLeaderboardRequest) (resp response.CrawlLeaderboardResponse)
}

type Facade struct {
	DBDAO      psql.DAOI
	CompDAO    compCrawl.DAOI
	CompMapper comp.MapperI
}

func (s Facade) CrawlLeaderboard(ctx context.Context, req request.CrawlLeaderboardRequest) (resp response.CrawlLeaderboardResponse) {
	compId := getCompId(req.CompName)
	url := s.CompMapper.BuildCrawlerUrl(compId)
	resp, err := s.CompDAO.CrawlLeaderboardData(url)
	if err != nil {
		log.Error(err)
		return response.CrawlLeaderboardResponse{
			Message: response.Message{
				ErrorLog: response.ErrorLogs{
					*err,
				},
			},
		}
	}
	for _, team := range resp.Teams {
		//TODO figure out if this should be 'update' instead of 'insert'
		exec := s.CompMapper.CreateInsertLeaderboardExec(compId, req.CompName, team)
		_, dbErr := s.DBDAO.InsertOne(ctx, exec)
		if err != nil {
			log.Error(dbErr)
			return response.CrawlLeaderboardResponse{
				Message: response.Message{
					ErrorLog: response.ErrorLogs{
						*dbErr,
					},
				},
			}
		}
	}
	//TODO add response mapping
	return resp
}

func getCompId(compName string) string {
	switch compName {
	case SixNations:
		return SixNationsId
	default:
		return ""
	}
}

const (
	SixNations   = "six nations"
	SixNationsId = "180659"
)
