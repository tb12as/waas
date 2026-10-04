package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tb12as/why-as-a-service/internal/app"
	"github.com/tb12as/why-as-a-service/internal/config"
	"github.com/tb12as/why-as-a-service/internal/database"
)

type APIReason struct {
	ID     uint
	Reason string
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	db, err := database.Load(cfg)
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		app.RandomReasonHandler(c, db)
	})
	r.Run(":" + cfg.AppPort)
}
