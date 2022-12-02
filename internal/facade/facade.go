package facade

import (
	"context"
	"github.com/calebtracey/rugby-crawler-api/internal/facade/comp"
	"github.com/calebtracey/rugby-models/pkg/dtos/leaderboard"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
)

//TODO create a response object to contain comp data + all other data

//go:generate mockgen -source=facade.go -destination=../mocks/mockFacade.go -package=mocks
type APIFacadeI interface {
	CrawlLeaderboardData(ctx context.Context, req leaderboard.Request) (resp leaderboard.Response)
	CrawlAllLeaderboardData(ctx context.Context) (resp leaderboard.Response)
}

type APIFacade struct {
	CompService comp.FacadeI
}

func (s APIFacade) CrawlLeaderboardData(ctx context.Context, req leaderboard.Request) (resp leaderboard.Response) {
	if validate := req.Validate(); validate != nil {
		return leaderboard.Response{
			Message: response.Message{ErrorLog: response.ErrorLogs{*validate}},
		}
	}
	resp = s.CompService.CrawlLeaderboard(ctx, req)

	return resp
}

func (s APIFacade) CrawlAllLeaderboardData(ctx context.Context) (resp leaderboard.Response) {
	resp = s.CompService.CrawlAllLeaderboards(ctx)

	return resp
}
