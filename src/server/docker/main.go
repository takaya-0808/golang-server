package main

import(	
	"github.com/gin-gonic/gin"
	"net/http"
)


type UserInfo struct {
	Username  string   `json:"username"`
	Email	  string   `json:"email"`
	Password  string   `json:"password"`
}

func main() {
	router := gin.Default()
	router.POST("/",Response)
	router.Run(":8008")
}

func Response(ctx *gin.Context) {

	userinfo := UserInfo{
		Username: "hoge",
		Email: "hhh@hhh.hhh",
		Password: "pass",
	}
	ctx.JSON(http.StatusOK, userinfo)
}