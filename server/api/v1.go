package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/wrfly/et/notify"
	"github.com/wrfly/et/server/asset"
	"github.com/wrfly/et/storage"
	"github.com/wrfly/et/types"
)

var pngFile, _ = asset.Asset("png/pixel.png")

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
	defer func() {
		c.Header("content-type", "image/png")
		c.Header("content-length", "126")
		c.Writer.Write(pngFile)
	}()

	var (
		taskID = c.Param("taskID")
		ip     = c.ClientIP()
		ua     = c.Request.UserAgent()
	)
	if len(taskID) > 40 {
		return
	}

	go func() {
		task, err := h.s.FindTask(taskID)
		if err != nil {
			logrus.Errorf("task [%s] not found", taskID)
			return
		}

		n := types.Notification{
			ID:     genNotificationID(fmt.Sprintf("%s-%s-%s", taskID, ip, ua)),
			TaskID: taskID,
			Event: types.OpenEvent{
				IP:   ip,
				UA:   ua,
				Time: time.Now(),
			},
		}

		// save notification
		if err := h.s.SaveNotification(n); err != nil {
			logrus.Errorf("save notification error: %s", err)
		}

		logrus.Debugf("send notification %+v", n)
		body, code, err := h.n.Send(task.NotifyTo, notify.NewContent(n, task.Comments))
		if err != nil {
			logrus.Errorf("send notification err: %s", err)
		}
		if code != http.StatusCreated {
			logrus.Errorf("handle task [%+v], body: %s, code: %d",
				task, body, code)
		}
	}()
}

func (h *Handler) Submit(c *gin.Context) {

	mailTo := c.Query("to")
	comments := c.Query("c")

	ID := genTaskID(mailTo, comments)
	t := types.Task{
		ID:       ID,
		NotifyTo: mailTo,
		Comments: comments,
	}
	logrus.Debugf("submit task %+v", t)
	if err := h.s.SaveTask(t); err != nil {
		logrus.Errorf("save task error: %s", err)
	}
	c.String(http.StatusOK,
		fmt.Sprintf("ID: %s\nhttps://track.kfd.me/t/%s\n",
			ID, ID))
}
