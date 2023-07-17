package config

import (
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"github.com/wrfly/ecp"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Listen  int  `default:"2020"`  // listen port
	Debug   bool `default:"false"` // log level
	Storage struct {
		Type string `default:"bolt"` // bolt/redis/mongo/...
		Bolt struct {
			Path string `default:"/data"`
		}
		Redis struct {
			Conn string
		}
		Mongo struct {
			Conn string
		}
	}
	SendGridAPI string
}

func (c *Config) Example() {
	cc := &Config{}
	if err := ecp.Parse(cc); err != nil {
		logrus.Fatalf("set default config value error: %s", err)
	}
	bs, _ := yaml.Marshal(cc)
	fmt.Printf("%s\n", bs)
	return
}

func (c *Config) Parse(filename string) error {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(bs, c)
}
