package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"rest/src/errs"
	"rest/src/infrastructure/gateways/rest/views"
	"rest/src/pkg/status"
	"time"
)

// ReadRecord godoc swagger
// @Summary Read record
// @Description API to get a record by id
// @Router /record/{id} [GET]
// @Tags Record
// @Accept json
// @Produce json
// @Param id path string true "record id"
// @Success 200 {object} views.R{data=views.RecordStruct}
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) ReadRecord(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		panic("id unparsed")
	}
	h.log.Info("keldi "+id)

	rand.Seed(time.Now().UnixNano())
	internalServer := rand.Intn(2)

	r, err := h.ctrl.ReadRecord(c, id)
	switch {
	case errors.Is(err, errs.ErrNotFound):
		h.handleNotFoundErr(c,http.StatusNotFound, status.ErrorNotFound)
	case internalServer == 0:
		h.handleInternalErr(c, http.StatusInternalServerError, status.ErrorInternalServer)
	case err == nil:
		h.handleSuccessResponse(c, views.Record(r))
	default:
		h.handleInternalErr(c, http.StatusInternalServerError, status.ErrorInternalServer)
	}
}