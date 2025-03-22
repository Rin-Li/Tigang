package routes

import (
	"Tigang/api"
	"Tigang/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("api")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})

		v1.POST("/users", api.UserRegister)
		v1.POST("/token", api.UserLogin)
		v1.POST("/users/reset_password_verify", api.ResetPasswordVerify)
		v1.PUT("/users/reset_password", api.ResetPassword)
		//Token
		auth := v1.Group("/")
		auth.Use(middleware.JWT())
		{
			//User
			auth.PUT("users/:id", api.UserUpdate)
			auth.GET("users/:id", api.ShowUser)

			//Record
			auth.PUT("records/:id", api.IncreaseRecord)
			auth.GET("records/:id", api.ShowListRecord)

		}

	}

	return r
}
