package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/wrfly/et/server/api"
	"github.com/wrfly/et/server/asset"
)

func Run(listen int, handler *api.Handler) error {
	e := gin.Default()
	// index
	for _, a := range asset.Data.List() {
		e.HEAD(a.Name(), gin.WrapH(asset.Data))
		e.GET(a.Name(), gin.WrapH(asset.Data))
	}

	// api
	apiG := e.Group("/api")
	{
		taskG := apiG.Group("/task")
		{
			taskG.POST("/submit", handler.SubmitTask)
			taskG.POST("/resume")
			taskG.GET("/get")
		}

		statusG := apiG.Group("/status")
		{
			statusG.GET("/task")
			statusG.GET("/notified")
		}
	}

	// track
	e.GET("/t/:taskID", handler.Open)

	return e.Run(fmt.Sprintf(":%d", listen))
}
