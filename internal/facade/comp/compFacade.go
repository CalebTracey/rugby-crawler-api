package comp

import (
	"context"
	"github.com/calebtracey/rugby-crawler-api/internal/dao/comp"
	"github.com/calebtracey/rugby-crawler-api/internal/dao/psql"
	"github.com/calebtracey/rugby-models/pkg/dtos/request"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	log "github.com/sirupsen/logrus"
	"strings"
)

//go:generate mockgen -destination=../../mocks/compmocks/mockFacade.go -package=compmocks . FacadeI
type FacadeI interface {
	CrawlLeaderboard(ctx context.Context, req request.LeaderboardRequest) (resp response.LeaderboardResponse)
}

type Facade struct {
	DbDAO    psql.DAOI
	CompDAO  comp.DAOI
	DbMapper psql.MapperI
}

func (s Facade) CrawlLeaderboard(ctx context.Context, req request.LeaderboardRequest) (resp response.LeaderboardResponse) {
	compId := getCompId(req.CompName)
	url := buildCrawlerUrl(compId)
	resp, err := s.CompDAO.CrawlLeaderboardData(url)
	if err != nil {
		log.Error(err)
		return response.LeaderboardResponse{
			Message: response.Message{
				ErrorLog: response.ErrorLogs{
					*err,
				},
			},
		}
	}
	for _, team := range resp.Teams {
		//TODO figure out if this should be 'update' instead of 'insert'
		exec := s.DbMapper.CreateInsertLeaderboardExec(compId, req.CompName, team)
		_, dbErr := s.DbDAO.InsertOne(ctx, exec)
		if dbErr != nil {
			log.Error(dbErr)
			return response.LeaderboardResponse{
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

func buildCrawlerUrl(compId string) string {
	return strings.Join([]string{CrawlBaseUrl, CrawlCompField, compId}, "")
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
	CrawlBaseUrl   = "https://www.espn.co.uk/rugby/"
	CrawlCompField = "table/_/league/"

	SixNations   = "six nations"
	SixNationsId = "180659"
)
