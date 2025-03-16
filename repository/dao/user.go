package dao

import (
	"Tigang/conf"
	"Tigang/model"
	"context"

	"gorm.io/gorm"
)

type UserDao struct{
    *gorm.DB
}

func NewUserDao (ctx context.Context) *UserDao {
    return &UserDao{conf.GetDB(ctx)}
}

func (dao *UserDao) ExistOrNotByUserName (userName string) (user *model.User, exit bool, err error){
    var count int64
    err = dao.DB.Model(&model.User{}).Where("username = ?", userName).Find(&user).Count(&count).Error
    if count == 0{
        return nil, false, err
    }
    return user, true, nil
}

func (dao *UserDao) CreateUser (user *model.User) error{
    return dao.DB.Create(user).Error
}

