package facade

import (
	"context"
	"github.com/calebtracey/rugby-crawler-api/internal/facade/comp"
	"github.com/calebtracey/rugby-crawler-api/internal/mocks/compmocks"
	"github.com/calebtracey/rugby-models/pkg/dtos"
	"github.com/calebtracey/rugby-models/pkg/dtos/leaderboard"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestAPIFacade_CrawlLeaderboardData(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockLeaderboardService := compmocks.NewMockFacadeI(ctrl)
	defer ctrl.Finish()

	type args struct {
		ctx context.Context
		req leaderboard.Request
	}
	tests := []struct {
		name        string
		CompService comp.FacadeI
		args        args
		wantResp    leaderboard.Response
	}{
		{
			name: "Happy Path",
			args: args{
				ctx: context.Background(),
				req: leaderboard.Request{
					Competitions: dtos.CompetitionList{
						{
							Name: "six nations",
						},
					},
				},
			},
			CompService: mockLeaderboardService,
			wantResp:    leaderboard.Response{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := APIFacade{
				CompService: tt.CompService,
			}
			mockLeaderboardService.EXPECT().CrawlLeaderboard(tt.args.ctx, tt.args.req).Return(tt.wantResp)
			if gotResp := s.CrawlLeaderboardData(tt.args.ctx, tt.args.req); !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("CrawlLeaderboardData() = %#v, want %#v", gotResp, tt.wantResp)
			}
		})
	}
}
