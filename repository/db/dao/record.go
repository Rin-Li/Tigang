package dao

import (
	"Tigang/conf"
	"Tigang/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type RecordDao struct{
    *gorm.DB
}

func NewRecordDao (ctx context.Context) *UserDao {
    return &UserDao{conf.GetDB(ctx)}
}


func (dao *RecordDao) UpdateRecord ()
