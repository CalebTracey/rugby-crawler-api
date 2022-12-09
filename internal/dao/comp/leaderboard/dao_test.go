package leaderboard

import (
	"github.com/calebtracey/rugby-crawler-api/internal/dao/comp/leaderboard/testdata"
	"github.com/calebtracey/rugby-models/pkg/dtos"
	"github.com/gocolly/colly"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestDAO_CrawlLeaderboardData(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/leaderboard", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(testdata.SixNationsLeaderboardHTML))
	})
	svr := httptest.NewServer(mux)

	testColly := colly.NewCollector()

	tests := []struct {
		name      string
		Collector *colly.Collector
		url       string
		wantResp  dtos.CompetitionLeaderboardData
		wantErr   bool
	}{
		{
			name:      "Happy Path",
			Collector: testColly,
			url:       svr.URL + "/leaderboard",
			wantResp: dtos.CompetitionLeaderboardData{
				Teams: dtos.TeamLeaderboardDataList{
					{
						Id:   "9",
						Name: "France",
						Abbr: "FRA",
						CompetitionStats: dtos.TeamCompetitionStats{
							GamesPlayed:       "5",
							WinCount:          "5",
							DrawCount:         "0",
							LossCount:         "0",
							Bye:               "0",
							PointsFor:         "141",
							PointsAgainst:     "73",
							TriesFor:          "17",
							TriesAgainst:      "7",
							BonusPointsTry:    "2",
							BonusPointsLosing: "0",
							BonusPoints:       "2",
							PointsDiff:        "+68",
							Points:            "25",
						},
					},
					{
						Id:   "3",
						Name: "Ireland",
						Abbr: "IRE",
						CompetitionStats: dtos.TeamCompetitionStats{
							GamesPlayed:       "5",
							WinCount:          "4",
							DrawCount:         "0",
							LossCount:         "1",
							Bye:               "0",
							PointsFor:         "168",
							PointsAgainst:     "63",
							TriesFor:          "24",
							TriesAgainst:      "4",
							BonusPointsTry:    "4",
							BonusPointsLosing: "1",
							BonusPoints:       "5",
							PointsDiff:        "+105",
							Points:            "21",
						},
					},
					{
						Id:   "1",
						Name: "England",
						Abbr: "ENG",
						CompetitionStats: dtos.TeamCompetitionStats{
							GamesPlayed:       "5",
							WinCount:          "2",
							DrawCount:         "0",
							LossCount:         "3",
							Bye:               "0",
							PointsFor:         "101",
							PointsAgainst:     "96",
							TriesFor:          "8",
							TriesAgainst:      "12",
							BonusPointsTry:    "1",
							BonusPointsLosing: "1",
							BonusPoints:       "2",
							PointsDiff:        "+5",
							Points:            "10",
						},
					},
					{
						Id:   "2",
						Name: "Scotland",
						Abbr: "SCOT",
						CompetitionStats: dtos.TeamCompetitionStats{
							GamesPlayed:       "5",
							WinCount:          "2",
							DrawCount:         "0",
							LossCount:         "3",
							Bye:               "0",
							PointsFor:         "92",
							PointsAgainst:     "121",
							TriesFor:          "11",
							TriesAgainst:      "15",
							BonusPointsTry:    "1",
							BonusPointsLosing: "1",
							BonusPoints:       "2",
							PointsDiff:        "-29",
							Points:            "10",
						},
					},
					{
						Id:   "4",
						Name: "Wales",
						Abbr: "WALES",
						CompetitionStats: dtos.TeamCompetitionStats{
							GamesPlayed:       "5",
							WinCount:          "1",
							DrawCount:         "0",
							LossCount:         "4",
							Bye:               "0",
							PointsFor:         "76",
							PointsAgainst:     "104",
							TriesFor:          "8",
							TriesAgainst:      "8",
							BonusPointsTry:    "0",
							BonusPointsLosing: "3",
							BonusPoints:       "3",
							PointsDiff:        "-28",
							Points:            "7",
						},
					},
					{
						Id:   "20",
						Name: "Italy",
						Abbr: "ITALY",
						CompetitionStats: dtos.TeamCompetitionStats{
							GamesPlayed:       "5",
							WinCount:          "1",
							DrawCount:         "0",
							LossCount:         "4",
							Bye:               "0",
							PointsFor:         "60",
							PointsAgainst:     "181",
							TriesFor:          "5",
							TriesAgainst:      "27",
							BonusPointsTry:    "0",
							BonusPointsLosing: "0",
							BonusPoints:       "0",
							PointsDiff:        "-121",
							Points:            "4",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DAO{
				Collector: tt.Collector,
			}
			gotResp, err := s.CrawlLeaderboardData(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("CrawlLeaderboardData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("CrawlLeaderboardData() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
