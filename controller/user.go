package controller

import (
	"Tigang/model"
	"Tigang/repository/dao"
	"Tigang/serializer"
	"context"
)

type UserService struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Email string `json:"email" form:"email"`
}
//Register
func (service *UserService) Register (ctx context.Context) serializer.Response{
	var user model.User
	
	userDao := dao.NewUserDao(ctx)

	_, exit, err := userDao.ExistOrNotByUserName(service.UserName)

	if err != nil{
		return serializer.Response{
			Status: 404,
			Msg:"database error",
			Error: err.Error(),
		}
	}

	if exit {
		return serializer.Response{
			Status: 404,
			Msg:"username is already exist",
		}
	}

	user = model.User{
		Username: service.UserName,
		PasswordDig: service.Password,
		Email: service.Email,
		TotalRecords: 0,
		ReminderInterval: 0,
	}

	err = userDao.CreateUser(&user)

	if err != nil{
		return serializer.Response{
			Status: 404,
			Msg:"database error",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Status: 200,
		Msg:"register success",
	}

}