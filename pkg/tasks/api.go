package tasks

import (
	"log"
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
		log.Println(ErrInvalidParameters.Wrap(err))
		generateResponse(c, Task{}, http.StatusBadRequest, ErrInvalidParameters)
		return
	}

	task := res.service.Get(c, uri.ID)
	generateResponse(c, task, http.StatusOK, nil)
}

func generateResponse(c *gin.Context, data Task, code int, err error) {
	if err != nil {
		c.JSON(code, gin.H{"data": responseResult{
			Task:  data,
			Error: err.Error(),
		}})
	} else {
		c.JSON(code, gin.H{"data": responseResult{
			Task: data,
		}})
	}
}
