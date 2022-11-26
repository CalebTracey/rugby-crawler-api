package psql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
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
		{
			name: "Sad Path - no rows",
			DB:   db,
			ctx:  context.Background(),
			exec: ``,
			wantErr: &response.ErrorLog{
				StatusCode: "404",
				RootCause:  "Not found in database",
			},
			mockErr:   sql.ErrNoRows,
			expectErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DAO{
				Db: tt.DB,
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
