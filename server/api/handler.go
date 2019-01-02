package api

import (
	"time"

	"github.com/wrfly/et/notify"
	"github.com/wrfly/et/server/asset"
	"github.com/wrfly/et/storage"
)

const (
	timeZone = "Asia/Shanghai"
)

var (
	pngFile  []byte
	local, _ = time.LoadLocation(timeZone)
)

func init() {
	_file, err := asset.Data.Asset("/png/pixel.png")
	if err != nil {
		panic(err)
	}
	pngFile = _file.Bytes()
}

type Handler struct {
	n notify.Notifier
	s storage.Database
}

func New(n notify.Notifier, s storage.Database) *Handler {
	return &Handler{n: n, s: s}
}
