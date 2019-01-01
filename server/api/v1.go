package api

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/wrfly/et/limiter"
	"github.com/wrfly/et/notify"
	"github.com/wrfly/et/server/asset"
	"github.com/wrfly/et/storage"
	"github.com/wrfly/et/types"
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

func (h *Handler) Open(c *gin.Context) {
	c.Writer.Write(pngFile)

	go func() {
		var (
			// trim suffix `.png` if found
			taskID = strings.TrimSuffix(c.Param("taskID"), ".png")
			ip     = c.ClientIP()
			ua     = c.Request.UserAgent()
		)
		if len(taskID) > 40 {
			return // basic check
		}

		task, err := h.s.FindTask(taskID)
		if err != nil {
			logrus.Warnf("find task [%s] error: %s", taskID, err)
			return
		}
		// check task state
		if task.BadState() {
			logrus.Warnf("task [%s] can not be processed due to bad state %s",
				task.ID, task.State)
			return
		}
		defer h.s.SaveTask(task) // save state
		if !checkTaskLimit(task) {
			return
		}

		// new notification
		notification := newNotification(*task, ip, ua)

		// save notification
		if err := h.s.SaveNotification(notification); err != nil {
			logrus.Errorf("save notification error: %s", err)
		}

		logrus.Debugf("send notification to %s", task.NotifyTo)
		body, code, err := h.n.Send(task.NotifyTo,
			notify.NewContent(notification, task.Comments))
		if err != nil {
			logrus.Errorf("send notification err: %s", err)
		}

		// sendgrid API returns 202
		if code != http.StatusAccepted {
			logrus.Errorf("handle task [%+v], body: %s, code: %d",
				task, body, code)
		}

		logrus.Debugf("send notification %+v done", notification)
	}()
}

func (h *Handler) SubmitTask(c *gin.Context) {
	r := taskRequest{}
	if err := c.BindJSON(&r); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, taskResponse{
			Error: err.Error(),
		})
		return
	}
	t := &types.Task{
		NotifyTo: r.NotifyTo,
		Comments: r.Comments,
		SubmitAt: time.Now(), // UTC
		Offset:   -time.Duration(r.Offset) * time.Minute,
	}

	// check ip
	if err := limiter.IP.Inc(c.ClientIP()); err != nil {
		logrus.Warnf("ip [%s] limiter: %s", c.ClientIP(), err)
		c.AbortWithStatusJSON(http.StatusBadRequest, taskResponse{
			Error: err.Error(),
		})
		return
	}

	if !validTask(t) {
		c.AbortWithStatusJSON(http.StatusBadRequest, taskResponse{
			Error: "bad request, check your email address and comments",
		})
		return
	}
	genTaskID(t)
	logrus.Debugf("submit task [%+v]", t)
	if err := h.s.SaveTask(t); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, taskResponse{
			Error: fmt.Sprintf("save task error: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, taskResponse{
		TaskID:    t.ID,
		TrackLink: fmt.Sprintf("%s/t/%s.png", DomainPrefix, t.ID),
	})
}
