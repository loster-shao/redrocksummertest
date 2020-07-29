package data

import (
	"github.com/jinzhu/gorm"
)

//用户
type User struct {
	gorm.Model
	User  string
	Pass  string
	Lv    int
	Root  bool
	Token string
}

//文件
type Down struct {
	gorm.Model
	Name     string//包名
	Personal bool  //是否私有
	User     string//上传用户
	Pass     string//密码
	Url      string//地址
}

//token
//头
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {// token里面添加用户信息，验证token后可能会用到用户信息
	Iss      string     `json:"iss"`
	Exp      string     `json:"exp"`
	IssueAt      string `json:"iat"`
	Username string     `json:"username"`
	Uid      int        `json:"id"`
}



