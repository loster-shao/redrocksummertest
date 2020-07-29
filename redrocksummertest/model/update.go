package model

import (
	"redrocksummertest/data"
)

func Update(file, password, url string, uid int, b bool) error {

	err := DB.Create(&data.Down{
		Name:     file,
		Personal: b,
		User:     string(uid),
		Pass:     password,
		Url:      url,
	}).Error
	return err
}
