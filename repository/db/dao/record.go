package dao

import (
	"Tigang/conf"
	"Tigang/repository/db/model"
	"context"
	"time"

	"gorm.io/gorm"
)

type RecordDao struct {
	*gorm.DB
}

func NewRecordDao(ctx context.Context) *RecordDao {
	return &RecordDao{conf.GetDB(ctx)}
}

func (dao *RecordDao) CreateRecord(record *model.Record) error {
	return dao.DB.Create(record).Error
}

func (dao *RecordDao) IncreaseRecordUser(uId uint) error {
	var user model.User
	err := dao.DB.Where("id = ?", uId).First(&user).Error
	if err != nil {
		return err
	}

	lastRecordDate := user.LastRecordTime.Truncate(24 * time.Hour)
	today := time.Now().Truncate(24 * time.Hour)

	if lastRecordDate.Equal(today.AddDate(0, 0, -1)) {
		user.ContinueRecordsDay += 1
	} else if lastRecordDate.Before(today) {
		user.ContinueRecordsDay = 1
	} 

	user.TotalRecords += 1
	user.LastRecordTime = time.Now()

	err = dao.DB.Save(&user).Error
	return err
}

func (dao *RecordDao) ShowRecord(uId uint) (model.User, error) {
	var user model.User
	err := dao.DB.Where("id = ?", uId).First(&user).Error
	return user, err
}

func (dao *RecordDao) ShowUserRecord(uId uint) ([]model.Record, error) {
	var records []model.Record
	err := dao.DB.Where("user_id = ?", uId).Find(&records).Error
	return records, err
}
