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
	"github.com/wrfly/et/types"
)

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

func (h *Handler) ResumeTask(c *gin.Context) {
	taskID := c.Query("id")
	if len(taskID) > 40 {
		return
	}
	task, err := h.s.FindTask(taskID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{
			"find task error": err.Error(),
		})
		return
	}
	if task.State != types.StateStopped {
		// hmm...
		return
	}
	task.State = types.StateResumed
	if err := h.s.SaveTask(task); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{
			"find notification error": err.Error(),
		})
	}

	// reset limiter
	limiter.Recv.Reset(task.NotifyTo)
	limiter.Sent.Reset(task.ID)
}

func (h *Handler) GetTask(c *gin.Context) {
	taskID := c.Query("id")
	if len(taskID) > 40 {
		return
	}
	task, err := h.s.FindTask(taskID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{
			"find task error": err.Error(),
		})
		return
	}
	ns, err := h.s.FindNotification(taskID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{
			"find notification error": err.Error(),
		})
		return
	}
	events := make([]types.OpenEvent, len(ns))
	for i, n := range ns {
		events[i] = n.Event
	}

	c.JSON(http.StatusOK, taskGetResponse{
		Comments: task.Comments,
		State:    task.State.String(),
		Submit:   task.SubmitAt,
		Events:   events,
	})
}
