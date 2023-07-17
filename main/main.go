package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/wrfly/ecp"

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
				for _, e := range ecp.List(conf, strings.ToUpper(appName)) {
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
			parser := ecp.New()
			parser.BuildKey = func(parentName, structName string, tag reflect.StructTag) string {
				return parentName + "_" + structName
			}
			if err := parser.Parse(conf, strings.ToUpper(appName)); err != nil {
				logrus.Fatalf("ecp parse error: %s", err)
			}

			if err := run(conf); err != nil {
				logrus.Fatal(err)
			}

			return nil
		},
	}

	app.Run(os.Args)
}
