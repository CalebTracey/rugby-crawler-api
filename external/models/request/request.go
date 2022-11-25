package request

type CrawlLeaderboardRequest struct {
	CompId   string `json:"compId,omitempty"`
	CompName string `json:"compName,omitempty"`
	Date     string `json:"date,omitempty"`
}
