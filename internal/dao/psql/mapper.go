package psql

import (
	"fmt"
	"github.com/calebtracey/rugby-models/pkg/dtos"
	log "github.com/sirupsen/logrus"
	"strconv"
)

//go:generate mockgen -source=mapper.go -destination=../../mocks/dbmocks/mockMapper.go -package=dbmocks
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
	s := td.CompetitionStats
	return fmt.Sprintf(UpdateLeaderboardExec,
		compId, compName, teamId, td.Name, td.Abbr, s.GamesPlayed, s.WinCount, s.DrawCount, s.LossCount,
		s.Bye, s.PointsFor, s.PointsAgainst, s.TriesFor, s.TriesAgainst, s.BonusPointsTry,
		s.BonusPointsLosing, s.BonusPoints, s.PointsDiff, s.Points, compId)
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
	s := td.CompetitionStats
	return fmt.Sprintf(InsertLeaderboardExec,
		compId, compName, teamId, td.Name, td.Abbr, s.GamesPlayed, s.WinCount, s.DrawCount, s.LossCount,
		s.Bye, s.PointsFor, s.PointsAgainst, s.TriesFor, s.TriesAgainst, s.BonusPointsTry,
		s.BonusPointsLosing, s.BonusPoints, s.PointsDiff, s.Points)
}

const (
	UpdateLeaderboardExec = `
		update comp_leaderboard
			set comp_id = '%v',
				comp_name = '%s',
				team_id = '%v',
				team_name = '%s',
				team_abbr = '%s',
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
			 team_abbr,
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
			values ('%v', '%s', '%v', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')
			on conflict (comp_id, team_id)
			do update
				set comp_id = excluded.comp_id,
					comp_name = excluded.comp_name,
					team_id = excluded.team_id,
					team_name = excluded.team_name,
					team_abbr = excluded.team_abbr,
					games_played = excluded.games_played,
					win_count = excluded.win_count,
					draw_count = excluded.draw_count,
					loss_count = excluded.loss_count,
					bye = excluded.bye,
					points_for = excluded.points_for,
					points_against = excluded.points_against,
					tries_for = excluded.tries_for,
					tries_against = excluded.tries_against,
					bonus_points_try = excluded.bonus_points_try,
					bonus_points_losing = excluded.bonus_points_losing,
					bonus_points = excluded.bonus_points,
					points_diff = excluded.points_diff,
					points = excluded.points_diff;`
)
