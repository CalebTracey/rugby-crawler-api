package facade

import (
	"context"
	"github.com/calebtracey/rugby-crawler-api/external/models/request"
	"github.com/calebtracey/rugby-crawler-api/external/models/response"
	"github.com/calebtracey/rugby-crawler-api/internal/facade/comp"
	"github.com/calebtracey/rugby-crawler-api/internal/mocks/compmocks"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestAPIFacade_CrawlLeaderboardData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLeaderboardService := compmocks.NewMockFacadeI(ctrl)
	type args struct {
		ctx context.Context
		req request.CrawlLeaderboardRequest
	}
	tests := []struct {
		name        string
		CompService comp.FacadeI
		args        args
		wantResp    response.CrawlLeaderboardResponse
	}{
		{
			name: "Happy Path",
			args: args{
				ctx: context.Background(),
				req: request.CrawlLeaderboardRequest{},
			},
			CompService: mockLeaderboardService,
			wantResp:    response.CrawlLeaderboardResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := APIFacade{
				CompService: tt.CompService,
			}
			mockLeaderboardService.EXPECT().CrawlLeaderboard(tt.args.ctx, tt.args.req).Return(tt.wantResp)
			if gotResp := s.CrawlLeaderboardData(tt.args.ctx, tt.args.req); !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("CrawlLeaderboardData() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
