package sqlstorage

import (
	"context"
	"database/sql"
	"errors"
	"rest/src/errs"
	"rest/src/model"
)

type recordRepo struct {
	s *Store
}

func (repo *recordRepo) GetRecords(ctx context.Context) ([]model.Record, error) {
	c, err := repo.GetAllRecords(ctx)
	if err != nil {
		return nil, err
	}
	records := make([]model.Record, len(c))
	for i, v := range c {
		records[i] = model.Record(v)
	}
	return records, nil
}

func (repo *recordRepo) GetRecord(ctx context.Context, id string) (model.RecordModel, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)

	q :=`SELECT * FROM records WHERE id = $1`

	var dbmodel dbRecord

	if err := sqlClient.QueryRowx(q, id).StructScan(&dbmodel); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.RecordModel{}, errs.Errf(errs.ErrNotFound, err.Error())
		}
		return model.RecordModel{}, err
	}

	return dbmodel.toModel(), nil
}

func (repo *recordRepo) GetAllRecords(ctx context.Context) ([]model.RecordModel, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)

	q :=`SELECT * FROM contact`

	rows := make([]dbRecord,0)
	if err :=sqlClient.Select(&rows,q); err != nil {
		return nil,err
	}

	contacts := make([]model.RecordModel, len(rows))
	for i := range rows {
		contacts[i] = rows[i].toModel()
	}

	return contacts, nil
}