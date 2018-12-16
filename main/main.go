package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/wrfly/ecp"
	"gopkg.in/urfave/cli.v2"

	"github.com/wrfly/et/config"
)

var (
	appName = "et"

	helpTemplate = `NAME:
    {{.Name}} - {{.Usage}}

AUTHOR:
    {{range .Authors}}{{ . }}
    {{end}}
VERSION:
    {{.Version}}

OPTIONS:
    {{range .VisibleFlags}}{{.}}
    {{end}}
`
)

func main() {

	app := &cli.App{
		Name:    appName,
		Usage:   "Self hosted email tracker.",
		Authors: author,
		Version: fmt.Sprintf("Version: %s\tCommit: %s\tDate: %s",
			Version, CommitID, BuildAt),
		CustomAppHelpTemplate: helpTemplate,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Usage:   "config file path",
				Aliases: []string{"c"},
			},
			&cli.BoolFlag{
				Name:    "example",
				Usage:   "config file example",
				Aliases: []string{"e"},
			},
			&cli.BoolFlag{
				Name:  "env-list",
				Usage: "config environment lists",
			},
		},
		Action: func(c *cli.Context) error {
			conf := &config.Config{}

			// pre actions
			if c.Bool("example") {
				conf.Example()
				return nil
			}
			if c.Bool("env-list") {
				for _, e := range ecp.List(conf, appName) {
					fmt.Println(e)
				}
				return nil
			}

			// parse config file
			if c.String("config") != "" {
				if err := conf.Parse(c.String("config")); err != nil {
					logrus.Fatalf("parse config file error: %s", err)
				}
			}

			// set default value
			if err := ecp.Parse(conf, appName); err != nil {
				logrus.Fatalf("ecp parse error: %s", err)
			}

			return run(conf)
		},
	}

	app.Run(os.Args)
}
