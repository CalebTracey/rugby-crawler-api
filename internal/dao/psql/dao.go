package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrInsertOne = errors.New("psql insert one error")
)

//go:generate mockgen -source=dao.go -destination=../../mocks/dbmocks/mockDao.go -package=dbmocks
type DAOI interface {
	InsertOne(ctx context.Context, exec string) (res sql.Result, err error)
}

type DAO struct {
	Db *sql.DB
}

func (s DAO) InsertOne(ctx context.Context, exec string) (resp sql.Result, err error) {
	resp, err = s.Db.ExecContext(ctx, exec)
	if err != nil {
		return resp, fmt.Errorf("%w: %s", ErrInsertOne, err)
	}
	return resp, nil
}
