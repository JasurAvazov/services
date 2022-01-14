package model

type Record interface {
	GetCode() string
	GetExternalCodes() map[string]string
}
