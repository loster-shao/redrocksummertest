package main

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

func main() {
	r := gin.Default()
	//初始化redis
	InitRedis()
	r.POST("/upload", func(c *gin.Context) {
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

		//关闭文件
		defer saveFile.Close()
	})

	r.Run(":8080")
}

//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"io"
//	"os"
//)
//

//func main()  {
//	fmt.Println("ok")
//	s := gin.Default()
//	s.GET("/test", func(c *gin.Context) {
//		file,_ := c.FormFile("file")
//
//		in, _  := file.Open()
//		defer in.Close()
//		out,_ := os.Create("./"+ file.Filename)
//		defer out.Close()
//		io.Copy(out, in)
//		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Filename))
//		c.File("./"+file.Filename)
//	})
//	s.Run(":8080")
//}

//func main()  {
//	model.DbConn()
//	model.DB.CreateTable(&data.Down{
//		Model:    gorm.Model{},
//		Name:     "",
//		Personal: false,
//		User:     "",
//		Pass:     "",
//		Url:      "",
//	})
//
//
//}

//func main() {
//	s := gin.Default()
//	s.POST("/test", func(c *gin.Context) {
//		ss := c.PostForm("u")
//		z, err := strconv.Atoi(ss)
//		fmt.Println(z)
//		fmt.Println(err)
//	})
//	s.Run(":8080")
//}

//func main() {
//
//	qrCode, _ := qr.Encode("niubi", qr.M, qr.Auto)
//	fmt.Println(qrCode)
//	qrCode, _ = barcode.Scale(qrCode, 256, 256)
//	fmt.Println(qrCode)
//	file, _ := os.Create(url)
//	fmt.Println(file)
//	defer file.Close()
//
//	png.Encode(file, qrCode)
//}
