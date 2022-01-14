package sqlstorage

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"rest/src/contracts"
	"rest/src/errs"
)

var TxContextKey = "SQL_TRANSACTION"

type dbSession struct {
	tx *sqlx.Tx
}

func (s *dbSession) Close(err error) {
	if err != nil {
		_ = s.tx.Rollback()
		return
	}
	_ = s.tx.Commit()
}

func (s *Store) StartSession(c context.Context) (contracts.Session, context.Context, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, nil, errs.Errf(errs.ErrSourceConnectionErr, err.Error())
	}
	c = context.WithValue(c, TxContextKey, tx)
	return &dbSession{tx}, c, nil
}

func (s *Store) sqlClientByCtx(ctx context.Context) sqlClient {
	if ctx == nil {
		return s.db
	}
	val := ctx.Value(TxContextKey)
	if val == nil {
		return s.db
	}
	tx, ok := val.(*sqlx.Tx)
	if !ok {
		return s.db
	}
	s.log.Debug("QUERIES INSIDE TRANSACTION")
	return tx
}

//sqlClient An interface to use for both sqlx.DB and sqlx.Tx (to use a transaction or not)
type sqlClient interface {
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
	DriverName() string
	Get(dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	MustExec(query string, args ...interface{}) sql.Result
	Exec(query string, args ...interface{}) (sql.Result, error)
	MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result
	NamedExec(query string, arg interface{}) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	Preparex(query string) (*sqlx.Stmt, error)
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	Rebind(query string) string
	Select(dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}