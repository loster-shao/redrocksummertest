package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"redrocksummertest/data"
	"strconv"
	"strings"
	"time"
)

func NewHeader() data.Header {
	return data.Header{
		Alg: "HS256",
		Typ: "JWT",
	}
}

func Create(username string, id int) string {
	header := NewHeader()
	payload := data.Payload{
		Iss:      "szs",                                                           //
		Exp:      strconv.FormatInt(time.Now().Add(10*time.Hour).Unix(), 10),//持续时间
		IssueAt:  strconv.FormatInt(time.Now().Unix(), 10),                  //签发时间
		Username: username,                                                        //用户名
		Uid:      id,                                                              //id
	}
	h, _ := json.Marshal(header)  //json初始化
	p, _ := json.Marshal(payload)
	headerBase64 := base64.StdEncoding.EncodeToString(h)
	payloadBase64 := base64.StdEncoding.EncodeToString(p)
	str1 := strings.Join([]string{headerBase64, payloadBase64}, ".")

	key := "szs"
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(str1))
	s := mac.Sum(nil)
	signature := base64.StdEncoding.EncodeToString(s)
	token := str1 + "." + signature
	return token
}
