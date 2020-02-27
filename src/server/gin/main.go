package main

import(
	"net/http"	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserInfo struct {
	Username  string   `json:"username"`
	Email	  string   `json:"email"`
	Password  string   `json:"password"`
}

type TokenResponse struct {
	AccessToken string  `json:"accesstoken"`
}

func main() {
	router := gin.Default()
	router.POST("/",Response)
	router.Run(":8800")
}

func Response(ctx *gin.Context) {

	var userinfo UserInfo
	ctx.BindJSON(&userinfo)
	token := CreatToken()
	tokenrequest := TokenResponse{
		AccessToken: *token,
	}
	ctx.JSON(http.StatusOK,tokenrequest)
}

func CreatToken() *string {
	u, err := uuid.NewRandom()
	if err != nil {
		panic(err)
		return nil
	}
	uu := u.String()
	return &uu
}