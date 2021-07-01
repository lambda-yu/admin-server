package models

import (
	"gorm.io/gorm"
)

type User struct {
	ModelBase
	UserName        string
	Password        string
	IsRoot          bool
	PermissionLevel int
}

func CreateUser(username string, password string, plevel int, isRoot bool) (err error) {
	return db.Create(&User{UserName: username, Password: password, IsRoot: isRoot, PermissionLevel: plevel}).Error
}

func GetUser(username, password string) User {
	var user User
	err := db.Where(&User{UserName: username, Password: password}).First(&user).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return User{ModelBase: ModelBase{ID: -1}}
	}
	return user
}

//func Login(username, password string) (bool, error){
//	user, err := GetUser(username, password)
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return false, err
//	}
//	if user.ID > 0{
//		return true, err
//	}
//	return false, nil
//}
