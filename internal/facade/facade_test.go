package facade

import (
	"context"
	"github.com/calebtracey/rugby-crawler-api/internal/facade/comp"
	"github.com/calebtracey/rugby-crawler-api/internal/mocks/compmocks"
	"github.com/calebtracey/rugby-models/request"
	"github.com/calebtracey/rugby-models/response"
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
		req request.LeaderboardRequest
	}
	tests := []struct {
		name        string
		CompService comp.FacadeI
		args        args
		wantResp    response.LeaderboardResponse
	}{
		{
			name: "Happy Path",
			args: args{
				ctx: context.Background(),
				req: request.LeaderboardRequest{},
			},
			CompService: mockLeaderboardService,
			wantResp:    response.LeaderboardResponse{},
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
