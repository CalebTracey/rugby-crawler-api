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
	CrawlLeaderboard(ctx context.Context, req request.CrawlLeaderboardRequest) (resp response.CrawlLeaderboardResponse)
}

type Facade struct {
	DBDAO      psql.DAOI
	CompDAO    compCrawl.DAOI
	CompMapper comp.MapperI
}

func (s Facade) CrawlLeaderboard(ctx context.Context, req request.CrawlLeaderboardRequest) (resp response.CrawlLeaderboardResponse) {
	//TODO create scrape url
	url := s.CompMapper.BuildCrawlerUrl(req.CompId)
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
	//compExec := s.CompMapper.MapAddPSQLCompetitionData(resp.CompId, resp.Name, resp.TeamIds)
	//_, dbErr := s.DBDAO.InsertOne(ctx, compExec)
	//if dbErr != nil {
	//	log.Error(dbErr)
	//	return response.CrawlLeaderboardResponse{
	//		Message: response.Message{
	//			ErrorLog: response.ErrorLogs{
	//				*dbErr,
	//			},
	//		},
	//	}
	//}
	//TODO add response mapping
	return resp
}
