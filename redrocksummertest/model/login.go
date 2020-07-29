package model

import (
	"fmt"
	"log"
	"redrocksummertest/data"
	"redrocksummertest/token"
	"strconv"
)

//登录
func Login(id, password string) (error, bool){
	var u data.User
	id0, err := strconv.Atoi(id)
	if err != nil {
		log.Println("账号格式有误:", err)
		err := fmt.Errorf("账号格式有误")//设置err内容
		return err, false
	}

	//查询id
	if err := DB.Where("id = ?", id0).First(&u).Error; err != nil{
		log.Println("数据库错误", err)
		err := fmt.Errorf("数据库错误")
		return err, false
	}

	//判断密码
	if password != u.Pass{
		log.Println()
		err := fmt.Errorf("密码错误")
		return err, false
	}

	//创建token
	token := token.Create(u.User, id0)
	user  := &data.User{
		User:  u.User,
		Pass:  u.Pass,
		Lv:    u.Lv,
		Root:  u.Root,
		Token: "",
	}
	DB.Model(&user).Update("token", token)
	return nil, true
}