package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"io"
	"os"
	"time"
)

var Client *redis.Client

func InitRedis() {
	fmt.Println("start")
	Client = redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "",
		DB:       0,
	})
}

func Upload(c *gin.Context) {
	//总上传字节
	total := 0
	//获得要上传的文件
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		panic(err)
	}

	//本地创建文件，如有直接打开，否则再创建
	saveFile, err := os.OpenFile(header.Filename, os.O_RDWR, os.ModePerm)
	if err != nil {
		saveFile, _ = os.Create(header.Filename)
	}

	//创建缓冲区，2048字节
	buf := make([]byte, 2048)

	//检测redis是否上传过，上传过就把存入的字节放入total
	if result, err := Client.Get(header.Filename).Int64(); err != nil {
		//覆盖total
		total = int(result)
		//设置文件seek 从文件开头开始偏移
		file.Seek(int64(total), io.SeekStart)

		fmt.Println("文件已上传过：", total)
	}

	//循环读取字节
	for {
		//read当前读取的字节
		read, err := file.Read(buf)
		if err != nil {
			//读取完毕，跳出
			break
		}

		fmt.Println(read)
		//跳过total字节追加保存
		saveFile.WriteAt(buf, int64(total))
		//将上传的字节数保存到redis
		total += read
		if err = Client.Set(header.Filename, total, time.Duration(0)).Err(); err != nil {
			fmt.Println("保存redis失败！")
		}
	}
	c.JSON(200,
		gin.H{
			"status":  200,
			"message": "上传成功",
		})

	//关闭文件
	defer saveFile.Close()
}
