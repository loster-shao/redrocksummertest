package router

import (
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
	"image/png"
	"os"
	"redrocksummertest/tool"

	"redrocksummertest/check"
	"redrocksummertest/model"
)

func Share(c *gin.Context) {
	share := c.PostForm("share_id")
	id := c.PostForm("users_id")

	url, err := model.Share(share)
	if check.CheckSql(share) || check.CheckSql(id) {
		tool.Sqlerr(c)
	}
	token := model.Token(id)
	uid, _, err := check.CheckToken(token)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "token验证错误",
		})
	}
	fmt.Println(uid)
	//生成二维码
	qrCode, _ := qr.Encode(url, qr.M, qr.Auto)

	qrCode, _ = barcode.Scale(qrCode, 256, 256)

	file, _ := os.Create(share+".png")
	defer file.Close()

	png.Encode(file, qrCode)


	c.JSON(200, gin.H{
		"status":  200,
		"message": url,
		"share":   share+".png",
	})
}
