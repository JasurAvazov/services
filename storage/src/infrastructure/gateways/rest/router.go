package rest

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func endpoints(r *gin.Engine, h *handler) {
	url := ginSwagger.URL("../swagger/doc.json")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.GET("/record/:id", h.ReadRecord)
}