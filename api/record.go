package api

import "github.com/gin-gonic/gin"


func UpdateRecord(c *gin.Context){
	id := c.Param("id")
	result := service.UpdateRecord(c, id)
	c.JSON(200, result)
}