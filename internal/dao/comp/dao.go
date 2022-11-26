package comp

import (
	"github.com/calebtracey/rugby-models/pkg/dtos"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"regexp"
)

//go:generate mockgen -destination=../../mocks/compmocks/mockDao.go -package=compmocks . DAOI
type DAOI interface {
	CrawlLeaderboardData(url string) (resp response.LeaderboardResponse, log *response.ErrorLog)
}

type DAO struct {
	Collector *colly.Collector
}

func (s DAO) CrawlLeaderboardData(url string) (resp response.LeaderboardResponse, errLog *response.ErrorLog) {
	s.Collector.OnHTML("table.standings > tbody", func(e *colly.HTMLElement) {
		re := regexp.MustCompile("[0-9]+")
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			href := el.ChildAttr("a.react-router-link", "href")
			team := dtos.TeamLeaderboardData{
				Id:   re.FindString(href),
				Name: el.ChildText("span.team-names"),
				Abbr: el.ChildText("abbr"),
				CompetitionStats: dtos.TeamCompetitionStats{
					GamesPlayed:       el.ChildText("td:nth-child(2)"),
					WinCount:          el.ChildText("td:nth-child(3)"),
					DrawCount:         el.ChildText("td:nth-child(4)"),
					LossCount:         el.ChildText("td:nth-child(5)"),
					Bye:               el.ChildText("td:nth-child(6)"),
					PointsFor:         el.ChildText("td:nth-child(7)"),
					PointsAgainst:     el.ChildText("td:nth-child(8)"),
					TriesFor:          el.ChildText("td:nth-child(9)"),
					TriesAgainst:      el.ChildText("td:nth-child(10)"),
					BonusPointsTry:    el.ChildText("td:nth-child(11)"),
					BonusPointsLosing: el.ChildText("td:nth-child(12)"),
					BonusPoints:       el.ChildText("td:nth-child(13)"),
					PointsDiff:        el.ChildText("td:nth-child(14)"),
					Points:            el.ChildText("td:nth-child(15)"),
				},
			}
			resp.LeaderboardData.Teams = append(resp.LeaderboardData.Teams, team)

		})
	})
	err := s.Collector.Visit(url)
	if err != nil {
		log.Error(err)
		errLog = mapError(err, url)
		return resp, errLog
	}

	s.Collector.Wait()
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
