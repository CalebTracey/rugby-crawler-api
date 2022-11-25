package models

type TeamList []Team

type Team struct {
	Id                string `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	GamesPlayed       string `json:"gamesPlayed,omitempty"`
	WinCount          string `json:"winCount,omitempty"`
	DrawCount         string `json:"drawCount,omitempty"`
	LossCount         string `json:"lossCount,omitempty"`
	Bye               string `json:"bye,omitempty"`
	PointsFor         string `json:"pointsFor,omitempty"`
	PointsAgainst     string `json:"pointsAgainst,omitempty"`
	TriesFor          string `json:"triesFor,omitempty"`
	TriesAgainst      string `json:"triesAgainst,omitempty"`
	BonusPointsTry    string `json:"bonusPointsTry,omitempty"`
	BonusPointsLosing string `json:"bonusPointsLosing,omitempty"`
	BonusPoints       string `json:"bonusPoints,omitempty"`
	PointsDiff        string `json:"pointsDiff,omitempty"`
	Points            string `json:"points,omitempty"`
}
