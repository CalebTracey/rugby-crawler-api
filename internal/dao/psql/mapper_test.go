package psql

import (
	"fmt"
	"github.com/calebtracey/rugby-models/pkg/dtos"
	"testing"
)

func TestMapper_UpdateLeaderboardExec(t *testing.T) {
	s := dtos.TeamCompetitionStats{
		GamesPlayed:       "1",
		WinCount:          "1",
		DrawCount:         "0",
		LossCount:         "0",
		Bye:               "0",
		PointsFor:         "32",
		PointsAgainst:     "12",
		TriesFor:          "2",
		TriesAgainst:      "1",
		BonusPointsTry:    "1",
		BonusPointsLosing: "0",
		BonusPoints:       "4",
		PointsDiff:        "+12",
		Points:            "16",
	}

	type args struct {
		compIdStr string
		compName  string
		td        dtos.TeamLeaderboardData
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Happy Path",
			args: args{
				compIdStr: "123",
				compName:  "Six Nations",
				td: dtos.TeamLeaderboardData{
					Id:   "1",
					Name: "England",
					Abbr: "ENG",
					CompetitionStats: dtos.TeamCompetitionStats{
						GamesPlayed:       "1",
						WinCount:          "1",
						DrawCount:         "0",
						LossCount:         "0",
						Bye:               "0",
						PointsFor:         "32",
						PointsAgainst:     "12",
						TriesFor:          "2",
						TriesAgainst:      "1",
						BonusPointsTry:    "1",
						BonusPointsLosing: "0",
						BonusPoints:       "4",
						PointsDiff:        "+12",
						Points:            "16",
					},
				},
			},
			want: fmt.Sprintf(UpdateLeaderboardExec,
				"123", "Six Nations", "1", "England", "ENG", s.GamesPlayed, s.WinCount, s.DrawCount, s.LossCount,
				s.Bye, s.PointsFor, s.PointsAgainst, s.TriesFor, s.TriesAgainst, s.BonusPointsTry,
				s.BonusPointsLosing, s.BonusPoints, s.PointsDiff, s.Points, "123"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Mapper{}
			if got := m.UpdateLeaderboardExec(tt.args.compIdStr, tt.args.compName, tt.args.td); got != tt.want {
				t.Errorf("UpdateLeaderboardExec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapper_InsertLeaderboardExec(t *testing.T) {
	s := dtos.TeamCompetitionStats{
		GamesPlayed:       "1",
		WinCount:          "1",
		DrawCount:         "0",
		LossCount:         "0",
		Bye:               "0",
		PointsFor:         "32",
		PointsAgainst:     "12",
		TriesFor:          "2",
		TriesAgainst:      "1",
		BonusPointsTry:    "1",
		BonusPointsLosing: "0",
		BonusPoints:       "4",
		PointsDiff:        "+12",
		Points:            "16",
	}

	type args struct {
		compIdStr string
		compName  string
		td        dtos.TeamLeaderboardData
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Happy Path",
			args: args{
				compIdStr: "123",
				compName:  "Six Nations",
				td: dtos.TeamLeaderboardData{
					Id:   "1",
					Name: "England",
					Abbr: "ENG",
					CompetitionStats: dtos.TeamCompetitionStats{
						GamesPlayed:       "1",
						WinCount:          "1",
						DrawCount:         "0",
						LossCount:         "0",
						Bye:               "0",
						PointsFor:         "32",
						PointsAgainst:     "12",
						TriesFor:          "2",
						TriesAgainst:      "1",
						BonusPointsTry:    "1",
						BonusPointsLosing: "0",
						BonusPoints:       "4",
						PointsDiff:        "+12",
						Points:            "16",
					},
				},
			},
			want: fmt.Sprintf(InsertLeaderboardExec,
				"123", "Six Nations", "1", "England", "ENG", s.GamesPlayed, s.WinCount, s.DrawCount, s.LossCount,
				s.Bye, s.PointsFor, s.PointsAgainst, s.TriesFor, s.TriesAgainst, s.BonusPointsTry,
				s.BonusPointsLosing, s.BonusPoints, s.PointsDiff, s.Points),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Mapper{}
			if got := m.InsertLeaderboardExec(tt.args.compIdStr, tt.args.compName, tt.args.td); got != tt.want {
				t.Errorf("InsertLeaderboardExec() = %v, want %v", got, tt.want)
			}
		})
	}
}
