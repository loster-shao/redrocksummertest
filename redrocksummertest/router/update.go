package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"redrocksummertest/tool"

	"redrocksummertest/check"
	"redrocksummertest/model"
)

func Update(c *gin.Context) {
	id := c.PostForm("id")
	token := model.Token(id)
	uid, _, err := check.CheckToken(token)
	if err != nil{
	}

	pw  := c.PostForm("password")
	b   := c.PostForm("bool")
	var Bool bool
	if b == "ok"{
		Bool = true
	}else {
		Bool = false
	}
	//读取文件
	file, err := c.FormFile("update")
	fmt.Println(file)
	if err != nil {
		log.Printf("file error:", err)
		c.JSON(300, gin.H{"status": 300, "message": "上传文件失败，请重新上传"})
		return
	}

	//保存本地
	f := fmt.Sprintf("./%s", file.Filename)
	if err := c.SaveUploadedFile(file, f); err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "上传文件失败，请重新上传",
		})
		return
	}

	url := "E:/GoProjects/src/redrocksummertest" + file.Filename
	if err := model.Update(file.Filename, pw, url, uid, Bool); err != nil{
		tool.DbErr(err, c)
		return
	}

	//发送成功信息，正在保存信息
	c.JSON(200,gin.H{
		"status":  200,
		"message": "恭喜上传成功",
	})
}


