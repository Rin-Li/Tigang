package model

import "gorm.io/gorm"


type Achievement struct {
	gorm.Model
	Name string `json:"name"`
	Description string `json:"description"`
	IconURL string `json:"icon_url"`

	Users []User `gorm:"many2many:user_achievements;"`
}