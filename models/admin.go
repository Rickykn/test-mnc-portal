package models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
