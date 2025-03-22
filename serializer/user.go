package serializer

import (
	"Tigang/repository/db/model"
	"time"
)

type User struct {
	ID uint `json:"id"`
	Username string `json:"username"`
	Record uint `json:"record"`
	LastRecordTime time.Time `json:"last_record_time"`
	ContinueRecordsDay uint `json:"continue_records_day"`
	ReminderInterval uint `json:"reminder_interval"`
}

func BuildUser(user model.User) *User {
	return &User{
		ID: user.ID,
		Username: user.Username,
		Record: user.TotalRecords,
		ReminderInterval: user.ReminderInterval,
		LastRecordTime: user.LastRecordTime,
		ContinueRecordsDay: user.ContinueRecordsDay,
	}
}