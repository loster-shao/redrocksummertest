package router

import (
	"github.com/gin-gonic/gin"
	"redrocksummertest/model"
)

func Find(c *gin.Context)  {
	find := c.Query("name")
	model.Find(find)
}
