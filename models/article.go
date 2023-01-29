package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID              int    `json:"id" gorm:"primary_key"`
	Content_article string `json:"content_article" gorm:"size:1000"`
	Source          string `json:"source"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
	User_id         *int `json:"user_id"`
	User            User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
