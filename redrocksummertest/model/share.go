package model

import (
	"fmt"
	"log"
	"redrocksummertest/data"
	"strconv"
)

func Share(id string) (string, error) {
	var down data.Down
	id0, err := strconv.Atoi(id)
	if err != nil{
		log.Println("数据格式不正确")
		return "", err
	}
	if err := DB.Where("id = ?", id0).First(&down).Error; err != nil{
		log.Println("数据库错误", err)
		err := fmt.Errorf("数据库错误")
		return "", err
	}
	return down.Url, nil
}
