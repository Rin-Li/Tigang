package service

import (
	"Tigang/pkg/util"
	"Tigang/repository/cache"
	"Tigang/repository/db/dao"
	"Tigang/repository/db/model"
	"Tigang/serializer"
	"context"
	"fmt"
	"strconv"
	"time"
)

type UserService struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Email string `json:"email" form:"email"`
}

type UserRestPasswordVerifyService struct {
	Email string `json:"email" form:"email"`
}

type UserRestPasswordService struct {
	Email string `json:"email" form:"email"`
	Code string `json:"code" form:"code"`
	Password string `json:"password" form:"password"`
}

type UserUpdateService struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Email string `json:"email" form:"email"`

}
//Register
func (service *UserService) Register (ctx context.Context) serializer.Response{
	var user model.User
	
	userDao := dao.NewUserDao(ctx)

	_, exit, err := userDao.ExistOrNotByUserEmail(service.Email)

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
			Msg:"user already registered",
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
//Login
func (service *UserService) Login(ctx context.Context) serializer.Response {
    var user *model.User
    userDao := dao.NewUserDao(ctx)
    user, exist, err := userDao.ExistOrNotByUserEmail(service.Email)

    if err != nil {
        return serializer.Response{
            Status: 404,
            Msg: "database error",
            Error: err.Error(),
        }
    }

	if !exist{
		return serializer.Response{
			Status: 404,
			Msg: "user not exist, please register",
		}
	}


    if user.PasswordDig != service.Password {
        return serializer.Response{
            Status: 404,
            Msg: "password error",
        }
    }

	token, _ := util.GenerateToken(user.ID)

    return serializer.Response{
        Status: 200,
        Msg: "login success",
        Data: serializer.TokenData{
			User: serializer.BuildUser(*user),
			Token: token,
		},
    }
}
//Rest Password
func (service *UserRestPasswordVerifyService) RestPasswordVerify(ctx context.Context) serializer.Response {
	code := util.GenerateCode()
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserEmail(service.Email)
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg: "database error",
			Error: err.Error(),
		}
	}


	if !exist {
		return serializer.Response{
			Status: 404,
			Msg: "user not exist",
		}
	}
	// put code into redis
	expiration := 60 * 5
	err = cache.SetResetCode(service.Email, code, time.Duration(expiration)*time.Second)

	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg: "set code failed",
			Error: err.Error(),
		}
	}
	// send email
	
	subject := "[Tigang App] Reset Password"
	body := fmt.Sprintf("Your reset password code is: %s", code)
	err = util.SendEmail(service.Email, subject, body)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg: "send email failed",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg: "send email success",}
}
//Reset Password
func (service *UserRestPasswordService) RestPassword(ctx context.Context) serializer.Response {
	//Verify code
	cacheCode, err := cache.GetResetCode(service.Email)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg: "get code failed",
			Error: err.Error(),
		}
	}

	if cacheCode != service.Code {
		return serializer.Response{
			Status: 404,
			Msg: "code error",
		}
	}

	//Update password
	userDao := dao.NewUserDao(ctx)
	err = userDao.UpdatePasswordByEmail(service.Email, service.Password)

	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg: "update password failed",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Status: 200,
		Msg: "update password success",
	}
}

func (service *UserUpdateService) Update(ctx context.Context, id string) serializer.Response{
	idInt, _ := strconv.Atoi(id)
	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserId(uint(idInt))
	if err != nil{
		return serializer.Response{
			Status: 404,
			Msg: "database error",
			Error: err.Error(),
		}
	}

	if !exist{
		return serializer.Response{
			Status: 404,
			Msg: "user not exist",
		}
	}

	if service.UserName != ""{
		user.Username = service.UserName
	}
	if service.Password != ""{
		user.PasswordDig = service.Password
	}
	if service.Email != ""{
		user.Email = service.Email
	}

	err = userDao.UpdateUserById(user, uint(idInt))
	
	if err != nil{
		return serializer.Response{
			Status: 404,
			Msg: "database error",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Status: 200,
		Msg: "update success",
	}

}

func GetUser(ctx context.Context, id string) serializer.Response{
	idInt, _ := strconv.Atoi(id)
	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserId(uint(idInt))
	if err != nil{
		return serializer.Response{
			Status: 404,
			Msg: "database error",
			Error: err.Error(),
		}
	}

	if !exist{
		return serializer.Response{
			Status: 404,
			Msg: "user not exist",
		}
	}

	return serializer.Response{
		Status: 200,
		Data: serializer.BuildUser(*user),
	}
}


