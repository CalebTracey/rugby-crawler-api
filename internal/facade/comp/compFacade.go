package comp

import (
	"context"
	"github.com/calebtracey/rugby-crawler-api/external/models/request"
	"github.com/calebtracey/rugby-crawler-api/external/models/response"
	"github.com/calebtracey/rugby-crawler-api/internal/dao/leaderboard"
	"github.com/calebtracey/rugby-crawler-api/internal/dao/psql"
	log "github.com/sirupsen/logrus"
	"strings"
)

//go:generate mockgen -destination=../../mocks/compmocks/mockFacade.go -package=compmocks . FacadeI
type FacadeI interface {
	CrawlLeaderboard(ctx context.Context, req request.CrawlLeaderboardRequest) (resp response.CrawlLeaderboardResponse)
}

type Facade struct {
	DbDAO             psql.DAOI
	LeaderboardDAO    leaderboard.DAOI
	LeaderboardMapper leaderboard.MapperI
}

func (s Facade) CrawlLeaderboard(ctx context.Context, req request.CrawlLeaderboardRequest) (resp response.CrawlLeaderboardResponse) {
	compId := getCompId(req.CompName)
	url := s.LeaderboardMapper.BuildCrawlerUrl(compId)
	resp, err := s.LeaderboardDAO.CrawlLeaderboardData(url)
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
		exec := s.LeaderboardMapper.CreateInsertLeaderboardExec(compId, req.CompName, team)
		_, dbErr := s.DbDAO.InsertOne(ctx, exec)
		if dbErr != nil {
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
	switch strings.ToLower(compName) {
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
