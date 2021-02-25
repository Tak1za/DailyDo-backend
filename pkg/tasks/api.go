package tasks

import (
	"net/http"

	"github.com/Tak1za/DailyDo-backend/internal/errors"
	"github.com/gin-gonic/gin"
)

var (
	// ErrInvalidParameters creates an invalid parameter error
	ErrInvalidParameters = errors.WrapError{Msg: "Invalid parameters"}
)

type resource struct {
	service Service
}

type uri struct {
	ID int `uri:"id" binding:"required"`
}

// RegisterHandlers registers handlers related to tasks
func RegisterHandlers(r *gin.RouterGroup, service Service) {
	res := resource{service}

	r.GET("/tasks/:id", res.get)
}

func (res resource) get(c *gin.Context) {
	var uri uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidParameters.Error()})
		return
	}

	task := res.service.Get(c, uri.ID)
	if (Task{}) == task {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": task})
}
