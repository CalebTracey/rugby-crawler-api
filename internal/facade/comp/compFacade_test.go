package comp

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/calebtracey/rugby-crawler-api/external/models/request"
	"github.com/calebtracey/rugby-crawler-api/external/models/response"
	"github.com/calebtracey/rugby-crawler-api/internal/dao/comp"
	"github.com/calebtracey/rugby-crawler-api/internal/dao/psql"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestFacade_CrawlLeaderboard(t *testing.T) {
	_, mock, _ := sqlmock.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDbDao := psql.NewMockDAOI(ctrl)
	mockCompDao := comp.NewMockDAOI(ctrl)

	type fields struct {
		DBDAO      psql.DAOI
		CompDAO    comp.DAOI
		CompMapper comp.MapperI
	}
	type args struct {
		ctx context.Context
		req request.CrawlLeaderboardRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp response.CrawlLeaderboardResponse
	}{
		{
			name: "Happy Path",
			fields: fields{
				DBDAO:      nil,
				CompDAO:    nil,
				CompMapper: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Facade{
				DBDAO:      tt.fields.DBDAO,
				CompDAO:    tt.fields.CompDAO,
				CompMapper: tt.fields.CompMapper,
			}
			if gotResp := s.CrawlLeaderboard(tt.args.ctx, tt.args.req); !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("CrawlLeaderboard() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
