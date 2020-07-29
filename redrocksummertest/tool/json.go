package tool

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Sqlerr(c *gin.Context)  {
	c.JSON(666,
		gin.H{
			"status": 666,
			"message": "sql危险！！！",
		})
}

func DbErr(err error, c *gin.Context)  {
	log.Println(time.Now(), "数据库错误: ", err)
	c.JSON(500,
		gin.H{
			"status" : 500,
			"数据库错误" : err,
		})
}

