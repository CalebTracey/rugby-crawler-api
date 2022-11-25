package psql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/calebtracey/rugby-crawler-api/external/models/response"
	log "github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

func TestDAO_InsertOne(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}(db)
	tests := []struct {
		name      string
		DB        *sql.DB
		ctx       context.Context
		exec      string
		wantResp  sql.Result
		wantErr   *response.ErrorLog
		mockErr   error
		expectErr bool
	}{
		{
			name:      "Happy Path",
			DB:        db,
			ctx:       context.Background(),
			exec:      ``,
			wantResp:  sqlmock.NewResult(int64(4), int64(12312123123)),
			expectErr: false,
		},
		{
			name: "Sad Path",
			DB:   db,
			ctx:  context.Background(),
			exec: ``,
			wantErr: &response.ErrorLog{
				StatusCode: "500",
				RootCause:  "error",
			},
			mockErr:   errors.New("error"),
			expectErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DAO{
				DB: tt.DB,
			}
			if !tt.expectErr {
				mock.ExpectExec(tt.exec).WillReturnResult(tt.wantResp)
			}
			if tt.expectErr {
				mock.ExpectExec(tt.exec).WillReturnResult(tt.wantResp).WillReturnError(tt.mockErr)
			}
			_, gotErr := s.InsertOne(tt.ctx, tt.exec)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("InsertOne() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

//func TestDAO_FindAll(t *testing.T) {
//	db, mock, _ := sqlmock.New()
//	defer func(db *sql.DB) {
//		err := db.Close()
//		if err != nil {
//			log.Error(err)
//		}
//	}(db)
//	type args struct {
//		ctx   context.Context
//		query string
//	}
//	tests := []struct {
//		name     string
//		DB       *sql.DB
//		args     args
//		mockCols []string
//		wantErr  *response.ErrorLog
//	}{
//		{
//			name: "Happy Path",
//			DB: db,
//			args: args{
//				ctx: context.Background(),
//				query: "{}",
//			},
//			mockCols: []string{""},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := DAO{
//				DB: tt.DB,
//			}
//			rows := sqlmock.NewRows(tt.mockCols).
//				AddRow(1, "comp 1", 10, "team 1").
//				AddRow(1, "comp 1", 5, "team 2")
//			mock.ExpectBegin()
//			gotRows, gotErr := s.FindAll(tt.args.ctx, tt.args.query)
//			if !reflect.DeepEqual(gotRows, tt.wantRows) {
//				t.Errorf("FindAll() gotRows = %v, want %v", gotRows, tt.wantRows)
//			}
//			if !reflect.DeepEqual(gotErr, tt.wantErr) {
//				t.Errorf("FindAll() gotErr = %v, want %v", gotErr, tt.wantErr)
//			}
//		})
//	}
//}
