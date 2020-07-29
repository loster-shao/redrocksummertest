package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"redrocksummertest/check"
	"redrocksummertest/model"
	"redrocksummertest/tool"
	"time"
)

func Login(c *gin.Context)  {
	id     := c.PostForm("users_id")
	pass   := c.PostForm("password")
	if check.CheckSql(id) || check.CheckSql(pass){
		tool.Sqlerr(c)
	}
	err, ok := model.Login(id, pass)
	if err != nil{
		tool.DbErr(err, c)
		return
	}

	if ok == false{
		log.Println(id, "用户名或密码输入错误，请重试",time.Now())
		c.JSON(300, gin.H{"status": 300, "message" : "用户名or密码输入错误,请重试！"})
	}
	c.JSON(200, gin.H{"status": 200, "message" : "恭喜你,登录成功!"})
}
