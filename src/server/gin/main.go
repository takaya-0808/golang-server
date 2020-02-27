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

	router.POST("/",Response)
	router.Run(":8800")
}

func Response(ctx *gin.Context) {

	userinfo := UserInfo{
		Username: "hoge",
		Email:    "1222@gmail.com",
		Password: "password",
	}

	ctx.JSON(http.StatusOK,userinfo)
}