package sqlstorage

import (
	"github.com/jmoiron/sqlx"
	"rest/src/pkg/logger"
	"rest/src/storage"
)

type Store struct {
	db          *sqlx.DB
	log         logger.Logger

	contact     *recordRepo
}

func New(db *sqlx.DB, log logger.Logger) *Store {
	return &Store{
		db:  db,
		log: log,
	}
}

func (s *Store) Record() storage.RecordRepository {
	if s.contact == nil {
		s.contact = &recordRepo{s}
	}

	return s.contact
}