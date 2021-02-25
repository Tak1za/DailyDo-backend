package tasks

import (
	"github.com/gin-gonic/gin"
)

// Service interface
type Service interface {
	Get(ctx *gin.Context, id int) Task
}

type service struct {
	repo Repository
}

// NewService function
func NewService(repo Repository) Service {
	return service{repo}
}

func (s service) Get(ctx *gin.Context, id int) Task {
	task := s.repo.Get(ctx, id)
	return task
}
