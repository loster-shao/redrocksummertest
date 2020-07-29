package model

import (
	"fmt"
	"log"
	"redrocksummertest/data"
	"strconv"
)

func Download(down_id string) (error, string) {
	id, err := strconv.Atoi(down_id)
	if err != nil{
		fmt.Println("error: ", err)
	}
	var down data.Down
	if err := DB.Where("id = ?", id).First(&down).Error; err != nil{
		log.Println("数据库错误", err)
		err := fmt.Errorf("数据库错误")
		return err, ""
	}
	return nil, down.Url
}
