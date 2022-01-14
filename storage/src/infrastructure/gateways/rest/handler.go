package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest/src/config"
	"rest/src/infrastructure/gateways/rest/views"
	controllers "rest/src/interfaces/api"
	"rest/src/pkg/logger"
	"rest/src/pkg/status"
	"rest/src/storage"
)

type handler struct {
	cfg     config.Config
	log     logger.Logger
	storage storage.Storage
	ctrl    *controllers.Controller
}

func newHandler(cfg config.Config, log logger.Logger, store storage.Storage) *handler {
	return &handler{
		cfg:     cfg,
		log:     log,
		storage: store,
		ctrl:    controllers.New(store, log),
	}
}

func (h *handler) handleSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, views.R{
		ErrorCode: status.NoError,
		Data:      data,
	})
}

func (h *handler) handleInternalErr(c *gin.Context, httpCode, errCode int,) {
	c.JSON(httpCode, views.R{
		ErrorCode: errCode,
	})
}

func (h *handler) handleNotFoundErr(c *gin.Context, httpCode, errCode int,) {
	c.JSON(httpCode, views.R{
		ErrorCode: errCode,
		Data:      nil,
	})
}