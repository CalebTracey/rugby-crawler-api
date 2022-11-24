package facade

import (
	"context"
	"github.com/calebtracey/rugby-crawler-api/external/models/request"
	"github.com/calebtracey/rugby-crawler-api/external/models/response"
	"github.com/calebtracey/rugby-crawler-api/internal/facade/comp"
)

type APIFacadeI interface {
	CompetitionCrawlData(ctx context.Context, req request.CompetitionCrawlRequest) (resp response.CompetitionCrawlResponse)
}

type APIFacade struct {
	CompService comp.FacadeI
}

func (s APIFacade) CompetitionCrawlData(ctx context.Context, req request.CompetitionCrawlRequest) (resp response.CompetitionCrawlResponse) {
	//TODO add validation
	resp = s.CompService.CompetitionCrawl(ctx, req)

	return resp
}
