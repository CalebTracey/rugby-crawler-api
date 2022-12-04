package comp

import (
	"context"
	"database/sql"
	"fmt"
	lb "github.com/calebtracey/rugby-crawler-api/internal/dao/comp/leaderboard"
	"github.com/calebtracey/rugby-crawler-api/internal/dao/psql"
	"github.com/calebtracey/rugby-models/pkg/dtos"
	"github.com/calebtracey/rugby-models/pkg/dtos/leaderboard"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

//go:generate mockgen -source=compFacade.go -destination=../../mocks/compmocks/mockFacade.go -package=compmocks
type FacadeI interface {
	CrawlLeaderboard(ctx context.Context, req leaderboard.Request) (resp leaderboard.Response)
	CrawlAllLeaderboards(ctx context.Context) (resp leaderboard.Response)
}

type Facade struct {
	DbDAO    psql.DAOI
	LBDao    lb.DAOI
	DbMapper psql.MapperI
}

func (s Facade) CrawlLeaderboard(ctx context.Context, req leaderboard.Request) (resp leaderboard.Response) {
	g, ctx := errgroup.WithContext(ctx)
	results := make([]dtos.CompetitionLeaderboardData, len(req.Competitions))

	for i, competition := range req.Competitions {
		i, competition := i, competition
		compName, compId := compId(competition.Name)

		g.Go(func() error {
			leaderboardData, err := crawlLeaderboard(ctx, compName, compId, s)
			if err == nil {
				results[i] = leaderboardData
			}
			return err
		})
	}

	if err := g.Wait(); err != nil {
		log.Error(err)
		resp.Message.ErrorLog = response.ErrorLogs{
			*mapError(err, fmt.Sprintf("%s", req)),
		}
	}
	resp.LeaderboardData = results

	return resp
}

func (s Facade) CrawlAllLeaderboards(ctx context.Context) (resp leaderboard.Response) {
	g, ctx := errgroup.WithContext(ctx)
	results := make([]dtos.CompetitionLeaderboardData, len(competitions))
	idx := 0

	for competitionName, competitionId := range competitions {
		competitionName, competitionId := competitionName, competitionId

		g.Go(func() error {
			leaderboardData, err := crawlLeaderboard(ctx, competitionName, competitionId, s)
			if err == nil {
				results[idx] = leaderboardData
			}
			idx++
			return err
		})
	}

	if err := g.Wait(); err != nil {
		log.Error(err)
		resp.Message.ErrorLog = response.ErrorLogs{
			*mapError(err, "all leaderboards query"),
		}
	}
	resp.LeaderboardData = results

	return resp
}

func crawlLeaderboard(ctx context.Context, name, id string, s Facade) (dtos.CompetitionLeaderboardData, error) {
	leaderboardData, err := s.LBDao.CrawlLeaderboardData(crawlerUrl(id))
	if err != nil {
		return dtos.CompetitionLeaderboardData{}, fmt.Errorf("error crawling leaderboard: %w", err)
	}

	leaderboardData.CompId = id
	leaderboardData.CompName = name

	for _, team := range leaderboardData.Teams {
		team := team
		if _, dbErr := s.DbDAO.InsertOne(ctx, s.DbMapper.InsertLeaderboardExec(id, name, team)); dbErr != nil {

			return dtos.CompetitionLeaderboardData{}, fmt.Errorf("error while inserting teams: %w", dbErr)
		}
	}
	return leaderboardData, nil
}

func mapError(err error, query string) (errLog *response.ErrorLog) {
	errLog = &response.ErrorLog{
		Query: query,
	}
	if err == sql.ErrNoRows {
		errLog.RootCause = "Not found in database"
		errLog.StatusCode = "404"
		return errLog
	}
	if err != nil {
		errLog.RootCause = err.Error()
	}
	errLog.StatusCode = "500"
	return errLog
}

func crawlerUrl(compId string) string {
	return strings.Join([]string{CrawlBaseUrl, CrawlCompField, compId}, "")
}

// compId returns competition name and id based on name match
func compId(compName string) (string, string) {
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
	competitions = map[string]string{
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
