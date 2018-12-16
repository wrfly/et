package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wrfly/et/server/api"
)

func Run(listen int, handler *api.Handler) error {
	e := gin.Default()

	// index
	// e.GET("/",asset.)

	// api
	apiG := e.Group("/api")
	{
		apiG.POST("/task/submit", handler.Submit)
	}

	// track
	e.GET("/t/:taskID", handler.Open)

	return e.Run(fmt.Sprintf(":%d", listen))
}
