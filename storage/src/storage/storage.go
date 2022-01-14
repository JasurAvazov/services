package storage

import (
	"context"
	"rest/src/contracts"
	"rest/src/model"
)

type Storage interface {
	contracts.ISessionProvider

	Record() RecordRepository
}

type RecordRetriever interface {
	GetRecords(ctx context.Context) ([]model.Record, error)
}

type RecordRepository interface {
	RecordRetriever
	GetRecord(context.Context, string) (model.RecordModel, error)
}