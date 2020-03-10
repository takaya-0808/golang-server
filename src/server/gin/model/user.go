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

func Connection() *gorm.DB {

	db,err := gorm.Open("mysql","root:password@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	return db
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
	if userinfo.ID != 0 {
		return &userinfo.ID
	}
	return nil
}