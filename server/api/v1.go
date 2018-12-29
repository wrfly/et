package api

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

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
	return &Handler{
		n: n,
		s: s,
	}
}

func (h *Handler) Open(c *gin.Context) {
	c.Writer.Write(pngFile)

	go func() {
		var (
			taskID = c.Param("taskID")
			ip     = c.ClientIP()
			ua     = c.Request.UserAgent()
		)
		// trim suffix `.png` if found
		taskID = strings.TrimSuffix(taskID, ".png")
		if len(taskID) > 40 {
			return
		}

		task, err := h.s.FindTask(taskID)
		if err != nil {
			logrus.Warnf("find task [%s] error: %s", taskID, err)
			return
		}
		if task.Opentimes > TaskLimit {
			logrus.Warnf("task [%s] open too many times", taskID)
			return
		}

		n := types.Notification{
			TaskID: taskID,
			Event: types.OpenEvent{
				IP:   ip,
				UA:   ua,
				Time: time.Now().Add(task.Adjust),
			},
		}
		n.ID = genNotificationID(n)

		// save notification
		if err := h.s.SaveNotification(n); err != nil {
			logrus.Errorf("save notification error: %s", err)
		}

		logrus.Debugf("send notification to %s", task.NotifyTo)
		body, code, err := h.n.Send(task.NotifyTo, notify.NewContent(n, task.Comments))
		if err != nil {
			logrus.Errorf("send notification err: %s", err)
		}

		// sendgrid service returns 202
		if code != http.StatusAccepted {
			logrus.Errorf("handle task [%+v], body: %s, code: %d",
				task, body, code)
		}

		logrus.Debugf("send notification %+v done", n)
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
	if r.LocalTime.IsZero() {
		r.LocalTime = time.Now().In(local)
	}

	t := types.Task{
		NotifyTo: r.NotifyTo,
		Comments: r.Comments,
		SubmitAt: r.LocalTime,
		Adjust: r.LocalTime.Sub(time.Now()).
			Truncate(time.Second),
	}

	if !validTask(t) {
		c.AbortWithStatusJSON(http.StatusBadRequest, taskResponse{
			Error: "bad request, check your email address and comments",
		})
		return
	}

	t.ID = genTaskID(t)
	logrus.Debugf("submit task [%s], notify to [%s] with comments [%s]",
		t.ID, t.NotifyTo, t.Comments)
	logrus.Debugf("submitAt: %s, adjust: %s", t.SubmitAt, t.Adjust)
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
