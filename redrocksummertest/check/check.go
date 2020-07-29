package check

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"redrocksummertest/data"
	"regexp"
	"strings"
)

//防sql注入（反正应该能挡）
func CheckSql(s string) bool {
	str := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	re, err := regexp.Compile(str)
	if err != nil {
		panic(err.Error())
		log.Println("error", err)
		return false
	}
	log.Println(re.MatchString(s))
	return re.MatchString(s)
}

//检测token
func CheckToken(token string) (uid int, username string, err error) {
	fmt.Println("token:", token)
	arr := strings.Split(token, ".")//切割
	if len(arr) != 3 {
		err = errors.New("token error1")//创建err为"oken error1"
		return
	}
	fmt.Printf("0%s\n 1%s\n 2%s\n", arr[0], arr[1], arr[2])


	_, err = base64.StdEncoding.DecodeString(arr[0])
	if err != nil {
		err = errors.New("token error2")
		return
	}
	pay, err := base64.StdEncoding.DecodeString(arr[1])
	if err != nil {
		err = errors.New("token error3")
		return
	}
	sign, err := base64.StdEncoding.DecodeString(arr[2])
	if err != nil {
		err = errors.New("token error4")
		return
	}

	str1 := arr[0] + "." + arr[1]

	key := []byte("szs")
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(str1))
	s := mac.Sum(nil)
	fmt.Println(sign)
	fmt.Println(s)
	if res := bytes.Compare(sign, s); res != 0 {//Compare 比较俩者是否相等
		fmt.Println("test")
		err = errors.New("token error5")
		return
	}

	var payload data.Payload
	if err := json.Unmarshal(pay,&payload); err != nil{
		log.Println("err:", err)
	}
	uid = payload.Uid
	username = payload.Username
	return
}