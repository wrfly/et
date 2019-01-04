package api

import (
	"html/template"
	"time"

	"github.com/wrfly/et/notify"
	"github.com/wrfly/et/server/asset"
	"github.com/wrfly/et/storage"
)

const (
	timeZone = "Asia/Shanghai"
)

var (
	pngFile    []byte
	local, _   = time.LoadLocation(timeZone)
	badgeGreen *template.Template
	badgeBlue  *template.Template
)

func init() {
	_file, err := asset.Data.Asset("/img/pixel.png")
	if err != nil {
		panic(err)
	}
	pngFile = _file.Bytes()

	_file, err = asset.Data.Asset("/img/blue.svg")
	if err != nil {
		panic(err)
	}
	badgeBlue, err = template.New("blue").
		Parse(string(_file.Bytes()))
	if err != nil {
		panic(err)
	}

	_file, err = asset.Data.Asset("/img/green.svg")
	if err != nil {
		panic(err)
	}
	badgeGreen, err = template.New("green").
		Parse(string(_file.Bytes()))
	if err != nil {
		panic(err)
	}
}

type Handler struct {
	n notify.Notifier
	s storage.Database
}

func New(n notify.Notifier, s storage.Database) *Handler {
	return &Handler{n: n, s: s}
}
