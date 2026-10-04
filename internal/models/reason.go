package models

import (
	"time"

	"gorm.io/gorm"
)

type Reason struct {
	ID        uint   `gorm:"primarykey"`
	Reason    string `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type APIReason struct {
	ID     *uint  `json:"id"`
	Reason string `json:"reason"`
}
