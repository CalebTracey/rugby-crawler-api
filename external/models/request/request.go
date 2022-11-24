package request

type CompetitionCrawlRequest struct {
	CompetitionID string `json:"competitionID,omitempty"`
	Date          string `json:"date,omitempty"`
}
