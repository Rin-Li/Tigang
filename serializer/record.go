package serializer

import "time"

type Record struct {
	ID uint `json:"id"`
	UserID uint `json:"user_id"`
	Time time.Time `json:"time"`
	TotalRecords uint `json:"total_records"`
}

func BuildRecord(record Record, totalRecords uint) Record{
	return Record{
		ID: record.ID,
		UserID: record.UserID,
		Time: record.Time,
		TotalRecords: totalRecords,
	}
}
