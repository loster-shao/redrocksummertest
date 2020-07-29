package model

import (
	"log"
	"redrocksummertest/data"
	"strconv"
)

func Token(uid string) string {
	id0, err := strconv.Atoi(uid)
	if err != nil{
		log.Println("err:", err)
		return ""
	}
	var u data.User
	if err := DB.Where("id = ?", id0).First(&u).Error; err != nil{
		log.Println("数据库错误", err)
	}
	return u.Token
}
