package psql

import (
	"fmt"
	"github.com/calebtracey/rugby-models/pkg/dtos"
	log "github.com/sirupsen/logrus"
	"strconv"
)

//go:generate mockgen -destination=../../mocks/dbmocks/mockMapper.go -package=dbmocks . MapperI
type MapperI interface {
	CreateUpdateLeaderboardExec(compIdStr, compName string, td dtos.TeamLeaderboardData) string
	CreateInsertLeaderboardExec(compIdStr, compName string, td dtos.TeamLeaderboardData) string
}

type Mapper struct{}

func (m Mapper) CreateUpdateLeaderboardExec(compIdStr, compName string, td dtos.TeamLeaderboardData) string {
	compId, err := strconv.Atoi(compIdStr)
	if err != nil {
		log.Error(err)
	}
	teamId, err := strconv.Atoi(td.Id)
	if err != nil {
		log.Error(err)
	}
	return fmt.Sprintf(UpdateLeaderboardExec,
		compId, compName, teamId, td.Name, td.GamesPlayed, td.WinCount, td.DrawCount, td.LossCount,
		td.Bye, td.PointsFor, td.PointsAgainst, td.TriesFor, td.TriesAgainst, td.BonusPointsTry,
		td.BonusPointsLosing, td.BonusPoints, td.PointsDiff, td.Points, compId)
}

func (m Mapper) CreateInsertLeaderboardExec(compIdStr, compName string, td dtos.TeamLeaderboardData) string {
	compId, err := strconv.Atoi(compIdStr)
	if err != nil {
		log.Error(err)
	}
	teamId, err := strconv.Atoi(td.Id)
	if err != nil {
		log.Error(err)
	}
	return fmt.Sprintf(InsertLeaderboardExec,
		compId, compName, teamId, td.Name, td.GamesPlayed, td.WinCount, td.DrawCount, td.LossCount,
		td.Bye, td.PointsFor, td.PointsAgainst, td.TriesFor, td.TriesAgainst, td.BonusPointsTry,
		td.BonusPointsLosing, td.BonusPoints, td.PointsDiff, td.Points)
}

const (
	UpdateLeaderboardExec = `
		update comp_leaderboard
			set comp_id = '%v',
				comp_name = '%s',
				team_id = '%v',
				team_name = '%s',
				games_played = '%s',
				win_count = '%s',
				draw_count = '%s',
				loss_count = '%s',
				bye = '%s',
				points_for = '%s',
				points_against = '%s',
				tries_for = '%s',
				tries_against = '%s',
				bonus_points_try = '%s',
				bonus_points_losing = '%s',
				bonus_points = '%s',
				points_diff = '%s',
				points = '%s'
			where comp_id = '%v';`

	InsertLeaderboardExec = `
		insert into comp_leaderboard
			(comp_id, 
			 comp_name, 
			 team_id, 
			 team_name, 
			 games_played, 
			 win_count, 
			 draw_count, 
			 loss_count, 
			 bye, 
			 points_for, 
			 points_against, 
			 tries_for, 
			 tries_against, 
			 bonus_points_try, 
			 bonus_points_losing, 
			 bonus_points, 
			 points_diff, 
			 points)
			values ('%v', '%s', '%v', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')
			returning comp_name;`
)
