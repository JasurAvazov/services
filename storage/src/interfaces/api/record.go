package controllers

import (
	"context"
	"rest/src/model"
)

func (ctrl *Controller) ReadRecord(ctx context.Context, id string) (model.RecordModel, error) {
	repo := ctrl.S.Record()
	records, err := repo.GetRecord(ctx, id)
	if err != nil {
		return model.RecordModel{}, err
	}
	return records, nil
}