package tasks

import (
	"github.com/Tak1za/DailyDo-backend/pkg/database"

	"github.com/gin-gonic/gin"
)

// Repository interface
type Repository interface {
	Get(ctx *gin.Context, id int) Task
}

type repository struct {
	db *database.DB
}

// NewRepository function
func NewRepository(db *database.DB) Repository {
	return repository{db}
}

func (r repository) Get(ctx *gin.Context, id int) Task {
	var ret Task
	_ = r.db.DB.First(&ret, id)
	return ret
}
