package models

import (
	"time"

	"gorm.io/gorm"
)

type Point struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Value_point int    `json:"value_point"`
	Status      string `json:"status"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
