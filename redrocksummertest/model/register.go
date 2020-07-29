package model

import "redrocksummertest/data"

//注册
func Register(username, password string) error {
	err := DB.Create(&data.User{
		User:  username,
		Pass:  password,
		Lv:    0,
		Root:  false,
	}).Error
	return err
}