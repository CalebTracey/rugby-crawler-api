package comp

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/calebtracey/rugby-crawler-api/internal/dao/comp"
	"github.com/calebtracey/rugby-crawler-api/internal/dao/psql"
	"github.com/calebtracey/rugby-crawler-api/internal/mocks/compmocks"
	"github.com/calebtracey/rugby-crawler-api/internal/mocks/dbmocks"
	"github.com/calebtracey/rugby-models/pkg/dtos"
	"github.com/calebtracey/rugby-models/pkg/dtos/request"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestFacade_CrawlLeaderboard(t *testing.T) {
	_, mock, _ := sqlmock.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDbDao := dbmocks.NewMockDAOI(ctrl)
	mockLeaderboardDAO := compmocks.NewMockDAOI(ctrl)
	mockDbMapper := dbmocks.NewMockMapperI(ctrl)
	type fields struct {
		DbDAO    psql.DAOI
		CompDAO  comp.DAOI
		DbMapper psql.MapperI
	}
	type args struct {
		ctx context.Context
		req request.LeaderboardRequest
	}
	tests := []struct {
		name                  string
		fields                fields
		args                  args
		exec                  string
		url                   string
		wantResp              response.LeaderboardResponse
		wantCrawlResp         response.LeaderboardResponse
		mockDbRes             sql.Result
		mockDbErr             *response.ErrorLog
		mockLeaderboardDAOErr *response.ErrorLog
		expectCrawlError      bool
		expectDbError         bool
	}{
		{
			name: "Happy Path",
			fields: fields{
				DbDAO:    mockDbDao,
				CompDAO:  mockLeaderboardDAO,
				DbMapper: mockDbMapper,
			},
			args: args{
				ctx: context.Background(),
				req: request.LeaderboardRequest{
					CompName: "Six Nations",
				},
			},
			exec: ``,
			url:  "https://www.espn.co.uk/rugby/table/_/league/180659",
			wantCrawlResp: response.LeaderboardResponse{
				Id:   "180659",
				Name: "Six Nations",
				Teams: dtos.TeamLeaderboardDataList{
					{},
				},
				Message: response.Message{},
			},
			wantResp: response.LeaderboardResponse{
				Id:   "180659",
				Name: "Six Nations",
				Teams: dtos.TeamLeaderboardDataList{
					{},
				},
				Message: response.Message{},
			},
			mockDbRes:             sqlmock.NewResult(int64(4), int64(12312123123)),
			mockDbErr:             nil,
			mockLeaderboardDAOErr: nil,
			expectCrawlError:      false,
			expectDbError:         false,
		},
		{
			name: "Sad Path - crawl error",
			fields: fields{
				DbDAO:    mockDbDao,
				CompDAO:  mockLeaderboardDAO,
				DbMapper: mockDbMapper,
			},
			args: args{
				ctx: context.Background(),
				req: request.LeaderboardRequest{
					CompName: "Six Nations",
				},
			},
			exec: ``,
			url:  "https://www.espn.co.uk/rugby/table/_/league/180659",
			wantCrawlResp: response.LeaderboardResponse{
				Message: response.Message{
					ErrorLog: response.ErrorLogs{
						{
							Query:      "https://test.url",
							RootCause:  "error",
							StatusCode: "500",
						},
					},
				},
			},
			wantResp: response.LeaderboardResponse{
				Message: response.Message{
					ErrorLog: response.ErrorLogs{
						{
							Query:      "https://test.url",
							RootCause:  "error",
							StatusCode: "500",
						},
					},
				},
			},
			mockDbRes: sqlmock.NewResult(int64(4), int64(12312123123)),
			mockDbErr: &response.ErrorLog{
				Query:      ``,
				RootCause:  "db error",
				StatusCode: "500",
			},
			mockLeaderboardDAOErr: &response.ErrorLog{
				Query:      "https://test.url",
				RootCause:  "error",
				StatusCode: "500",
			},
			expectCrawlError: true,
			expectDbError:    false,
		},
		{
			name: "Sad Path - database error",
			fields: fields{
				DbDAO:    mockDbDao,
				CompDAO:  mockLeaderboardDAO,
				DbMapper: mockDbMapper,
			},
			args: args{
				ctx: context.Background(),
				req: request.LeaderboardRequest{
					CompName: "Six Nations",
				},
			},
			exec: ``,
			url:  "https://www.espn.co.uk/rugby/table/_/league/180659",
			wantCrawlResp: response.LeaderboardResponse{
				Id:   "180659",
				Name: "Six Nations",
				Teams: dtos.TeamLeaderboardDataList{
					{},
				},
				Message: response.Message{},
			},
			wantResp: response.LeaderboardResponse{
				Message: response.Message{
					ErrorLog: response.ErrorLogs{
						{
							Query:      ``,
							RootCause:  "db error",
							StatusCode: "500",
						},
					},
				},
			},
			mockDbRes: sqlmock.NewResult(int64(4), int64(12312123123)),
			mockDbErr: &response.ErrorLog{
				Query:      ``,
				RootCause:  "db error",
				StatusCode: "500",
			},
			mockLeaderboardDAOErr: &response.ErrorLog{
				Query:      "https://test.url",
				RootCause:  "error",
				StatusCode: "500",
			},
			expectCrawlError: false,
			expectDbError:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Facade{
				DbDAO:    tt.fields.DbDAO,
				CompDAO:  tt.fields.CompDAO,
				DbMapper: tt.fields.DbMapper,
			}
			mock.ExpectBegin()
			mockLeaderboardDAO.EXPECT().CrawlLeaderboardData(tt.url).
				DoAndReturn(func(url string) (response.LeaderboardResponse, *response.ErrorLog) {
					if tt.expectCrawlError {
						return tt.wantCrawlResp, tt.mockLeaderboardDAOErr
					}
					return tt.wantCrawlResp, nil
				})
			if !tt.expectCrawlError {
				mockDbMapper.EXPECT().CreateInsertLeaderboardExec(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.exec)
				mockDbDao.EXPECT().InsertOne(tt.args.ctx, tt.exec).
					DoAndReturn(func(ctx context.Context, exec string) (sql.Result, *response.ErrorLog) {
						if tt.expectDbError {
							mock.ExpectExec(tt.exec).WillReturnError(errors.New("db error"))
							return sqlmock.NewErrorResult(errors.New("db error")), tt.mockDbErr
						}
						mock.ExpectExec(tt.exec)
						return tt.mockDbRes, nil
					})
			}

			if gotResp := s.CrawlLeaderboard(tt.args.ctx, tt.args.req); !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("CrawlLeaderboard() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
