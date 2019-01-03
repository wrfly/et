package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) StatusTask(c *gin.Context) {
	c.JSON(http.StatusOK, h.s.Status())
}

func (h *Handler) StatusNotified(c *gin.Context) {
	c.JSON(http.StatusOK, h.s.Status())
}
