package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func DbConn()  {

	userName := "root"
	password := ""
	host     := "127.0.0.1"
	dbName   := "summer"
	port     := 3306
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", userName,password, host, port, dbName )
	db, err := gorm.Open("mysql", connArgs)
	//fmt.Println("db", db)
	if err != nil {
		fmt.Println("err:", err)
	}
	db.SingularTable(true)
	DB = db
}

