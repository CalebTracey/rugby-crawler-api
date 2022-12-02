package comp

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/calebtracey/rugby-crawler-api/internal/dao/comp"
	"github.com/calebtracey/rugby-crawler-api/internal/dao/psql"
	"github.com/calebtracey/rugby-crawler-api/internal/mocks/compmocks"
	"github.com/calebtracey/rugby-crawler-api/internal/mocks/dbmocks"
	"github.com/calebtracey/rugby-models/pkg/dtos"
	"github.com/calebtracey/rugby-models/pkg/dtos/leaderboard"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestFacade_CrawlLeaderboard(t *testing.T) {
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
		req leaderboard.Request
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		exec             string
		url              string
		wantResp         leaderboard.Response
		wantCrawlResp    dtos.CompetitionLeaderboardData
		wantCrawlError   error
		mockDbRes        sql.Result
		mockDbErr        error
		expectCrawlError bool
		expectDbError    bool
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
				req: leaderboard.Request{
					Competitions: dtos.CompetitionList{
						{
							Name: "six nations",
						},
					},
				},
			},
			exec: ``,
			url:  "https://www.espn.co.uk/rugby/table/_/league/180659",
			wantCrawlResp: dtos.CompetitionLeaderboardData{
				CompId:   SixNationsId,
				CompName: SixNations,
				Teams: dtos.TeamLeaderboardDataList{
					{
						Id:   "1",
						Name: "Team 1",
					},
					{
						Id:   "2",
						Name: "Team 2",
					},
				},
			},
			wantCrawlError: nil,
			wantResp: leaderboard.Response{
				LeaderboardData: dtos.CompetitionLeaderboardDataList{
					{
						CompId:   SixNationsId,
						CompName: SixNations,
						Teams: dtos.TeamLeaderboardDataList{
							{
								Id:   "1",
								Name: "Team 1",
							},
							{
								Id:   "2",
								Name: "Team 2",
							},
						},
					},
				},
				Message: response.Message{},
			},
			mockDbRes:        sqlmock.NewResult(int64(4), int64(12312123123)),
			mockDbErr:        nil,
			expectCrawlError: false,
			expectDbError:    false,
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
				req: leaderboard.Request{
					Competitions: dtos.CompetitionList{
						{
							Name: "six nations",
						},
					},
				},
			},
			exec:          ``,
			url:           "https://www.espn.co.uk/rugby/table/_/league/180659",
			wantCrawlResp: dtos.CompetitionLeaderboardData{},
			wantResp: leaderboard.Response{
				LeaderboardData: dtos.CompetitionLeaderboardDataList{
					{
						CompId:   "",
						CompName: "",
						Teams:    dtos.TeamLeaderboardDataList(nil),
					},
				},
				Message: response.Message{
					ErrorLog: response.ErrorLogs{
						{
							Query: fmt.Sprintf("%s", leaderboard.Request{
								Competitions: dtos.CompetitionList{
									{
										Name: "six nations",
									},
								},
							}),
							RootCause:  "error crawling leaderboard: error",
							StatusCode: "500",
						},
					},
				},
			},
			mockDbRes:        sqlmock.NewResult(int64(4), int64(12312123123)),
			mockDbErr:        errors.New("db error"),
			wantCrawlError:   errors.New("error"),
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
				req: leaderboard.Request{
					Competitions: dtos.CompetitionList{
						{
							Name: "six nations",
						},
					},
				},
			},
			exec: ``,
			url:  "https://www.espn.co.uk/rugby/table/_/league/180659",
			wantCrawlResp: dtos.CompetitionLeaderboardData{
				CompId:   SixNationsId,
				CompName: SixNations,
				Teams: dtos.TeamLeaderboardDataList{
					{
						Id:   "1",
						Name: "Team 1",
					},
					{
						Id:   "2",
						Name: "Team 2",
					},
				},
			},
			wantResp: leaderboard.Response{
				LeaderboardData: dtos.CompetitionLeaderboardDataList{
					{
						CompId:   "",
						CompName: "",
						Teams:    dtos.TeamLeaderboardDataList(nil),
					},
				},
				Message: response.Message{
					ErrorLog: response.ErrorLogs{
						{
							Query: fmt.Sprintf("%s", leaderboard.Request{
								Competitions: dtos.CompetitionList{
									{
										Name: "six nations",
									},
								},
							}),
							RootCause:  "error crawling leaderboard: error",
							StatusCode: "500",
						},
					},
				},
			},
			mockDbRes:        sqlmock.NewResult(int64(4), int64(12312123123)),
			mockDbErr:        errors.New("db error"),
			wantCrawlError:   errors.New("error"),
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
			mockLeaderboardDAO.EXPECT().CrawlLeaderboardData(tt.url).Return(tt.wantCrawlResp, tt.wantCrawlError)
			if !tt.expectCrawlError {
				mockDbMapper.EXPECT().CreateInsertLeaderboardExec(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(tt.exec).MaxTimes(len(tt.wantCrawlResp.Teams))
				mockDbDao.EXPECT().InsertOne(gomock.Any(), tt.exec).
					DoAndReturn(func(ctx context.Context, exec string) (sql.Result, error) {
						if tt.expectDbError {
							return sqlmock.NewErrorResult(errors.New("db error")), tt.mockDbErr
						}
						return tt.mockDbRes, nil
					}).MaxTimes(len(tt.wantCrawlResp.Teams))
			}

			if gotResp := s.CrawlLeaderboard(tt.args.ctx, tt.args.req); !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("CrawlLeaderboard() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
