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
	for _, a := range asset.List() {
		e.HEAD(a.Name(), gin.WrapF(asset.Handler))
		e.GET(a.Name(), gin.WrapF(asset.Handler))
	}

	// api
	apiG := e.Group("/api")
	{
		taskG := apiG.Group("/task")
		{
			taskG.POST("/submit", handler.SubmitTask)
			taskG.POST("/resume", handler.ResumeTask)
			taskG.GET("/get", handler.GetTask)
		}

		statusG := apiG.Group("/status")
		{
			statusG.GET("/task.svg", handler.StatusTask)
			statusG.GET("/notified.svg", handler.StatusNotified)
		}
	}

	// track
	e.GET("/t/:taskID", handler.Open)

	return e.Run(fmt.Sprintf(":%d", listen))
}
