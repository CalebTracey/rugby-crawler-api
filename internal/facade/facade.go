package facade

import (
	"context"
	"github.com/calebtracey/rugby-crawler-api/internal/facade/comp"
	"github.com/calebtracey/rugby-models/pkg/dtos/request"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
)

//TODO create a response object to contain comp data + all other data

//go:generate mockgen -destination=../mocks/mockFacade.go -package=mocks . APIFacadeI
type APIFacadeI interface {
	CrawlLeaderboardData(ctx context.Context, req request.LeaderboardRequest) (resp response.LeaderboardResponse)
	CrawlAllLeaderboardData(ctx context.Context) (resp response.AllLeaderboardsResponse)
}

type APIFacade struct {
	CompService comp.FacadeI
}

func (s APIFacade) CrawlLeaderboardData(ctx context.Context, req request.LeaderboardRequest) (resp response.LeaderboardResponse) {
	//TODO add request validation
	resp = s.CompService.CrawlLeaderboard(ctx, req)

	return resp
}

func (s APIFacade) CrawlAllLeaderboardData(ctx context.Context) (resp response.AllLeaderboardsResponse) {
	//TODO add request validation
	resp = s.CompService.CrawlAllLeaderboards(ctx)
	//TODO add response mapping func
	return resp
}
