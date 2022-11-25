package request

type CrawlLeaderboardRequest struct {
	CompId string `json:"compId,omitempty"`
	Date   string `json:"date,omitempty"`
}
