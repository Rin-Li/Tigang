package service

import (
	"Tigang/serializer"
	"strconv"

	"github.com/gin-gonic/gin"
)


func UpdateRecord(c *gin.Context, id string) serializer.Response{
	uId := strconv.Atoi(id)
	recordDao := dao.NewRecordDao(ctx)
}