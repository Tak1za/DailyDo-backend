package main

import (
	"log"
	"net/http"

	"github.com/Tak1za/DailyDo-backend/pkg/database"
	"github.com/Tak1za/DailyDo-backend/pkg/tasks"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	log.Println("Application started")

	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	dbAccess := database.Initiate(db)

	hs := &http.Server{
		Addr:    ":8080",
		Handler: buildHandler(dbAccess),
	}

	if err := hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}

func buildHandler(db *database.DB) http.Handler {
	r := gin.Default()

	r.Use(cors.Default())

	rg := r.Group("/api/v1")

	tasks.RegisterHandlers(rg, tasks.NewService(tasks.NewRepository(db)))

	return r
}
