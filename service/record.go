package service

import (
	"Tigang/repository/db/dao"
	"Tigang/repository/db/model"
	"Tigang/serializer"
	"context"
	"strconv"
	"time"
)


func IncreaseRecord(ctx context.Context, id string) serializer.Response{
	uId, _ := strconv.Atoi(id)

	recordDao := dao.NewRecordDao(ctx)

	err := recordDao.IncreaseRecordUser(uint(uId))

	if err != nil{
		return serializer.Response{
			Status: 404,
			Msg: "database error",
			Error: err.Error(),
		}
	}

	record := model.Record{
		UserId: uint(uId),
		Time: time.Now(),
	}

	err = recordDao.CreateRecord(&record)

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

func ShowListRecord(ctx context.Context, id string) serializer.Response{
	uId, _ := strconv.Atoi(id)

	recordDao := dao.NewRecordDao(ctx)

	records, err := recordDao.ShowUserRecord(uint(uId))

	if err != nil{
		return serializer.Response{
			Status: 404,
			Msg: "database error",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Status: 200, 
		Data: serializer.BuildListRecords(records),	
	}
}




