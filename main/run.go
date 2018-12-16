package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wrfly/et/storage/bolt"

	"github.com/wrfly/et/config"
	"github.com/wrfly/et/notify"
	"github.com/wrfly/et/server"
	"github.com/wrfly/et/server/api"
)

func run(c *config.Config) error {
	// debug log
	if c.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	n := notify.NewSendgridNotifier(c.SendGridAPI)
	s, err := bolt.New(c.Storage.Bolt.Path)
	if err != nil {
		return err
	}

	return server.Run(c.Listen, api.New(n, s))
}
