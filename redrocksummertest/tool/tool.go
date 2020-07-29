package tool

import (
	"fmt"
	"github.com/skip2/go-qrcode"
)

//生成二维码
var png []byte
func Erweima() {
	err := qrcode.WriteFile("127.0.0.1:8080/szs", qrcode.Medium, 256, "qr.png")
	if err != nil {
		fmt.Println("write error")
	}
}