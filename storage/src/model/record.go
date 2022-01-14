package model

type RecordModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

func (n RecordModel) GetCode() string {
	return ""
}

func (n RecordModel) GetExternalCodes() map[string]string {
	return nil
}