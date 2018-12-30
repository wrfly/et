package config

import (
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/sirupsen/logrus"

	"gopkg.in/yaml.v2"

	"github.com/wrfly/ecp"
)

func init() {
	ecp.GetKey = func(parentName, structName string, tag reflect.StructTag) string {
		return parentName + "_" + structName
	}
}

type Config struct {
	Listen  int  `default:"2020"`  // listen port
	Debug   bool `default:"false"` // log level
	Storage struct {
		Type string `default:"bolt"` // bolt/redis/mongo/...
		Bolt struct {
			Path string `default:"/tmp"`
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
	if err := ecp.Default(cc); err != nil {
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
