package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tb12as/why-as-a-service/internal/app"
	"github.com/tb12as/why-as-a-service/internal/config"
	"github.com/tb12as/why-as-a-service/internal/database"
	"github.com/tb12as/why-as-a-service/internal/middleware"
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

	var r *gin.Engine
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	} else {
		r = gin.Default() // local only
	}
	r.Use(middleware.RateLimiter())

	rh := app.ReasonHandler{DB: db}
	r.GET("/", rh.RandomReasonHandler)

	r.Run(":" + cfg.AppPort)
}
