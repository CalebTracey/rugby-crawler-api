package psql

import (
	"context"
	"database/sql"
	"github.com/calebtracey/rugby-models/response"
	log "github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=../../mocks/dbmocks/mockDao.go -package=dbmocks . DAOI
type DAOI interface {
	InsertOne(ctx context.Context, exec string) (res sql.Result, error *response.ErrorLog)
}

type DAO struct {
	Db *sql.DB
}

func (s DAO) InsertOne(ctx context.Context, exec string) (resp sql.Result, err *response.ErrorLog) {
	resp, sqlErr := s.Db.ExecContext(ctx, exec)
	if sqlErr != nil {
		log.Error(sqlErr)
		err = mapError(sqlErr, exec)
		return resp, err
	}
	return resp, nil
}

func mapError(err error, query string) (errLog *response.ErrorLog) {
	errLog = &response.ErrorLog{
		Query: query,
	}
	if err == sql.ErrNoRows {
		errLog.RootCause = "Not found in database"
		errLog.StatusCode = "404"
		return errLog
	}
	if err != nil {
		errLog.RootCause = err.Error()
	}
	errLog.StatusCode = "500"
	return errLog
}
