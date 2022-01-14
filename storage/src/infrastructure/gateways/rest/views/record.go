package views

import "rest/src/model"

type RecordStruct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

func Record(model model.RecordModel) RecordStruct {
	return RecordStruct{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
	}
}