package api

import (
	"Tigang/service"

	"github.com/gin-gonic/gin"
)


func IncreaseRecord(c *gin.Context){
	id := c.Param("id")
	result := service.IncreaseRecord(c, id)
	c.JSON(200, result)
}

func ShowListRecord(c *gin.Context){
	id := c.Param("id")
	result := service.ShowListRecord(c, id)
	c.JSON(200, result)
}