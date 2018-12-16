package server

import (
	"github.com/gin-gonic/gin"
	"github.com/wrfly/et/notify"
)

type Server struct {
	e *gin.Engine
	n notify.Notifier
	// storage

}

func New(listen int) (*Server, error) {
	return nil, nil
}
