package controllers

import (
	"rest/src/pkg/logger"
	"rest/src/storage"
)

type Controller struct {
	S storage.Storage
	logger.Logger
}

func New(store storage.Storage, logger logger.Logger) *Controller {
	return &Controller{S: store, Logger: logger}
}