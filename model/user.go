package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	PasswordDig string `json:"password"`
	Email    string `json:"email"`
	TotalRecords int `json:"total_records"`
	ReminderInterval int `json:"reminder_interval"`
}


func (user *User) SetPassword(password string) error{
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil{
		return err
	}
	user.PasswordDig= string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) error{
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDig), []byte(password))
	if err != nil{
		return err
	}
	return nil
}