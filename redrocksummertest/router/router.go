package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(app *gin.Engine)  {
	app.LoadHTMLFiles("./index.html")//添加模板
	InitRedis()
	app.POST("/register", Register)//注册OK    ce
	app.POST("/login", Login)//登录OK          ce
	app.GET("/download", Download)//下载       有点问题
	app.POST("/update", Update)//更新          ce
	app.POST("upload", Upload)//上传           ce
	app.POST("/share", Share)//分享
	app.GET("/join", Find)//寻找（未完工，其实也不难）
	app.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK,"index.html",nil)
	})
}
