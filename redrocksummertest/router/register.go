package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"redrocksummertest/check"
	"redrocksummertest/model"
	"redrocksummertest/tool"
)

func Register(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if check.CheckSql(username) || check.CheckSql(password){
		tool.Sqlerr(c)
	}

	if username == "" || password == ""{
		fmt.Println("账号名或密码不能为空")
		c.JSON(300, gin.H{"status": 300, "message" : "用户名或密码不能为空"})
		return
	}

	if err := model.Register(username, password); err != nil{
		tool.DbErr(err, c)
		return
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "注册成功",
	})
}
