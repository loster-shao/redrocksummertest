package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

type Reader struct {
	io.Reader      //读取大小等
	Total   int64  //总共多大
	Current int64  //现在多大
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	r.Current += int64(n)
	fmt.Printf("\r当前进度%.2f%", float64(r.Current*10000/r.Total)/100)
	return
}

func Download(c *gin.Context)  {
	file := c.Query("file")
	//err, url := model.Download(file)
	//if err != nil {
	//	fmt.Println("error:", err)
	//}
	url := "E:/GoProjects/src/redrocksummertes" +
		"t/444.jpg"
	DownloadFileProgress(url, file)
}

func DownloadFileProgress(url, filename string) {
	r, err := http.Get(url)
	if err != nil{
		panic(err)
	}
	defer func() {
		_ = r.Body.Close()
	}()
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("err", err)
	}
	defer func() {
		_ = f.Close()
	}()

	reader := &Reader{
		Reader:  r.Body,
		Total:   r.ContentLength,
	}
	n, err := io.Copy(f, reader)//
	fmt.Println(n, err)
}
