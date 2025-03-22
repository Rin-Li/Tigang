package model

import (
	"time"
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	UserId uint `json:"user_id"`
	Time time.Time `json:"time"`
}