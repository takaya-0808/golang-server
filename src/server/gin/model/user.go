package model

import (
	"github.com/jinzhu/gorm"
   _ "github.com/go-sql-driver/mysql"
)

type UserInfo struct {
	ID        int      `json:"id"`
	Username  string   `json:"username"`
	Email	  string   `json:"email"`
	Password  string   `json:"password"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&UserInfo{})
}

func NewInsert(name string, email string, password string, db *gorm.DB) {

	userinfo := UserInfo{Username:name, Email:email, Password:password}
	db.Create(&userinfo)
}

func LoginProcess(name string, password string, db *gorm.DB) *int {

	var userinfo UserInfo
	db.Where("Username = ? AND Password = ?", name, password).Find(&userinfo)
	if &userinfo != nil {
		return &userinfo.ID
	}
	return nil
}