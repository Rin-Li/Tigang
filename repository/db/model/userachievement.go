package model

import (
	"time"

	"gorm.io/gorm"
)

type UserAchievement struct {
	gorm.Model
	UserID uint `json:"user_id"`
	AchievementID uint `json:"achievement_id"`
	UnlockedAt time.Time `json:"unlocked_at"`
}