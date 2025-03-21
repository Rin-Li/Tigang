package api

import (
	"Tigang/service"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var UserRegister service.UserService
	if err := c.ShouldBind(&UserRegister); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	} else {
		c.JSON(200, UserRegister.Register(c))
	}
}

func UserLogin(c *gin.Context) {
	var UserLogin service.UserService
	if err := c.ShouldBind(&UserLogin); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	} else {
		c.JSON(200, UserLogin.Login(c))
	}
}

func ResetPasswordVerify(c *gin.Context) {
	var ResetPasswordVerify service.UserRestPasswordVerifyService
	if err := c.ShouldBind(&ResetPasswordVerify); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	} else {
		c.JSON(200, ResetPasswordVerify.RestPasswordVerify(c))
	}
}

func ResetPassword(c *gin.Context) {
	var ResetPassword service.UserRestPasswordService
	if err := c.ShouldBind(&ResetPassword); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	} else {
		c.JSON(200, ResetPassword.RestPassword(c))
	}
}

func UserUpdate(c *gin.Context) {
	var UserUpdate service.UserUpdateService
	if err := c.ShouldBind(&UserUpdate); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	} else {
		c.JSON(200, UserUpdate.Update(c, c.Param("id")))
	}
}

func ShowUser(c *gin.Context) {
	id := c.Param("id")
	result := service.ShowUser(c, id)
	c.JSON(result.Status, result)
}
