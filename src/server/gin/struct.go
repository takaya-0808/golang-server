package main

import(
	"github.com/jinzhu/gorm"
)

type StrcutUserInfo struct {
	ID        int      `json:"id"`
	Username  string   `json:"username"`
	Email	  string   `json:"email"`
	Password  string   `json:"password"`
}

func New(db *gorm.DB) *StrcutUserInfo {
	return  &StrcutUserInfo{db: db}
}
