package comp

import (
	"context"
	"github.com/calebtracey/rugby-crawler-api/internal/dao/comp"
	"github.com/calebtracey/rugby-crawler-api/internal/dao/psql"
	"github.com/calebtracey/rugby-models/pkg/dtos"
	"github.com/calebtracey/rugby-models/pkg/dtos/request"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	log "github.com/sirupsen/logrus"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

//go:generate mockgen -destination=../../mocks/compmocks/mockFacade.go -package=compmocks . FacadeI
type FacadeI interface {
	CrawlLeaderboard(ctx context.Context, req request.LeaderboardRequest) (resp response.LeaderboardResponse)
	CrawlAllLeaderboards(ctx context.Context) (resp response.AllLeaderboardsResponse)
}

type Facade struct {
	DbDAO    psql.DAOI
	CompDAO  comp.DAOI
	DbMapper psql.MapperI
}

func (s Facade) CrawlLeaderboard(ctx context.Context, req request.LeaderboardRequest) (resp response.LeaderboardResponse) {
	compName, compId := getCompId(req.CompName)
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
	resp.LeaderboardData.CompName = compName
	resp.LeaderboardData.CompId = compId
	//TODO make this concurrent
	for _, team := range resp.LeaderboardData.Teams {
		exec := s.DbMapper.CreateInsertLeaderboardExec(compId, compName, team)
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
	//TODO make better response struct and add mapping
	return resp
}

func (s Facade) CrawlAllLeaderboards(ctx context.Context) (resp response.AllLeaderboardsResponse) {
	var req request.LeaderboardRequest
	//TODO make this concurrent
	for name, id := range CompMap {
		req.CompName = name
		req.CompId = id
		lb := s.CrawlLeaderboard(ctx, req)
		//TODO make a mapping func for this
		resp.LeaderboardDataList = append(resp.LeaderboardDataList, dtos.CompetitionLeaderboardData{
			CompId:   id,
			CompName: name,
			Teams:    lb.LeaderboardData.Teams,
		})
	}
	return resp
}

func buildCrawlerUrl(compId string) string {
	return strings.Join([]string{CrawlBaseUrl, CrawlCompField, compId}, "")
}

func getCompId(compName string) (string, string) {
	c := cases.Title(language.English)
	switch c.String(compName) {
	case SixNations:
		return SixNations, SixNationsId
	case RugbyWorldCup:
		return RugbyWorldCup, RugbyWorldCupId
	case Premiership:
		return Premiership, PremiershipId
	case Top14:
		return Top14, Top14Id
	case UnitedRugbyChampionship:
		return UnitedRugbyChampionship, UnitedRugbyChampionshipId
	case RugbyChampionship:
		return RugbyChampionship, RugbyChampionshipId
	default:
		return "", ""
	}
}

var (
	CompMap = map[string]string{
		SixNations:              SixNationsId,
		RugbyWorldCup:           RugbyWorldCupId,
		Premiership:             PremiershipId,
		Top14:                   Top14Id,
		UnitedRugbyChampionship: UnitedRugbyChampionshipId,
		RugbyChampionship:       RugbyChampionshipId,
	}
)

const (
	CrawlBaseUrl   = "https://www.espn.co.uk/rugby/"
	CrawlCompField = "table/_/league/"

	SixNations   = "Six Nations"
	SixNationsId = "180659"

	RugbyWorldCup   = "Rugby World Cup"
	RugbyWorldCupId = "164205"

	Premiership   = "Premiership"
	PremiershipId = "267979"

	Top14   = "Top 14"
	Top14Id = "270559"

	UnitedRugbyChampionship   = "United Rugby Championship"
	UnitedRugbyChampionshipId = "270557"

	RugbyChampionship   = "Rugby Championship"
	RugbyChampionshipId = "244293"
)
