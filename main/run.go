package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/wrfly/et/config"
)

func run(c *config.Config) error {
	// debug log
	if c.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	return nil
}
