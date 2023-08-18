package model

import (
	"github.com/jinzhu/gorm"
	//"github.com/RaymondCode/simple-demo/dao"
	"simple-demo-main/dao"
)

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth checks if authentication information exists
func CheckAuth(username, password string) (bool, error, int) {
	var auth Auth
	err := dao.DB.Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err, -1
	}

	if auth.ID > 0 {
		return true, nil, auth.ID
	}

	return false, nil, -1
}
