package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tb12as/why-as-a-service/internal/models"
	"github.com/tb12as/why-as-a-service/pkg"
	"gorm.io/gorm"
)

func getRandomReason(db *gorm.DB) (*models.APIReason, error) {
	random := models.APIReason{}

	db.Model(models.Reason{}).Limit(1).Order("rand()").Find(&random)
	if random.ID == nil {
		return nil, fmt.Errorf("No data")
	}

	return &random, nil
}

func RandomReasonHandler(c *gin.Context, db *gorm.DB) {
	r, err := getRandomReason(db)
	if err != nil {
		pkg.Fail(c, err.Error())
		return
	}

	pkg.Ok(c, r)
}
