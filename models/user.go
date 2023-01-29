package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           int    `json:"id" gorm:"primary_key"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Point_reward int    `json:"point_reward"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
