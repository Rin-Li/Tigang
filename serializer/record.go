package serializer

import (
	"Tigang/repository/db/model"
	"time"
)

type Record struct {
	ID uint `json:"id"`
	Time time.Time `json:"time"`
	TotalRecords uint `json:"total_records"`
}

func BuildRecord(record model.Record) Record{
	return Record{
		ID: record.ID,
		Time: record.Time,
	}
}

func BuildListRecords(records []model.Record) []Record{
	var recordList []Record
	for _, record := range records{
		recordList = append(recordList, BuildRecord(record))
	}
	return recordList
}