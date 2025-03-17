package model

import (
	"time"
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	UserID uint `json:"user_id"`
	Time time.Time `json:"time"`
	Note string `json:"note,omitempty"`
}