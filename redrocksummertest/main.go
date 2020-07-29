package main

import (
	"github.com/gin-gonic/gin"
	"redrocksummertest/model"

	"redrocksummertest/router"
)

func main()  {
	app := gin.Default()
	model.DbConn()
	router.SetupRouter(app)
	app.Run(":8080")
}
