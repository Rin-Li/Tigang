package serializer

import "Tigang/repository/db/model"

type User struct {
	ID uint `json:"id"`
	Username string `json:"username"`
	Record int `json:"record"`
	ReminderInterval int `json:"reminder_interval"`
}

func BuildUser(user model.User) *User {
	return &User{
		ID: user.ID,
		Username: user.Username,
		Record: user.TotalRecords,
		ReminderInterval: user.ReminderInterval,
	}
}