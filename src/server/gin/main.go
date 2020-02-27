package main

import(
	"net/http"	
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username  string   `json:"username"`
	Email	  string   `json:"email"`
	Password  string   `json:"password"`
}

func main() {
	router := gin.Default()

	userinfo := UserInfo{
		Username: "hoge",
		Email:    "1222@gmail.com",
		Password: "password",
	}

	router.POST("/",func(ctx *gin.Context){
		ctx.JSON(http.StatusOK,userinfo)
	})
	router.Run(":8800")
}