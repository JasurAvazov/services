package sqlstorage

import (
	"rest/src/model"
)

type dbRecord struct {
	ID        int    `db:"id"`
	Name string `db:"name"`
	Description string `db:"description"`
}

func (raw *dbRecord) toModel() model.RecordModel {
	result := model.RecordModel{
		ID:          raw.ID,
		Name:        raw.Name,
		Description: raw.Description,
	}
	return result
}