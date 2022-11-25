package facade

import (
	"context"
	"github.com/calebtracey/rugby-crawler-api/external/models/request"
	"github.com/calebtracey/rugby-crawler-api/external/models/response"
	"github.com/calebtracey/rugby-crawler-api/internal/facade/comp"
)

//TODO create response object to contain leaderboard data + all other data

type APIFacadeI interface {
	CompetitionCrawlData(ctx context.Context, req request.CrawlLeaderboardRequest) (resp response.CrawlLeaderboardResponse)
}

type APIFacade struct {
	CompService comp.FacadeI
}

func (s APIFacade) CompetitionCrawlData(ctx context.Context, req request.CrawlLeaderboardRequest) (resp response.CrawlLeaderboardResponse) {
	//TODO add validation
	resp = s.CompService.CrawlLeaderboard(ctx, req)

	return resp
}
