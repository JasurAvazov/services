package rest

import (
	"github.com/gin-gonic/gin"
	"rest/src/config"
	"rest/src/pkg/logger"
	"rest/src/storage"
)

// NewAPI ...
func NewAPI(cfg config.Config, log logger.Logger, r *gin.Engine, store storage.Storage) *gin.Engine {
	h := newHandler(cfg, log, store)
	endpoints(r, h)

	return r
}