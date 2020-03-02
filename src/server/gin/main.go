package main

import(
	"net/http"	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gin/model"
)

type Error struct {
	Error  	  string   `json:"error"`    
}

type LoginForm struct {
	Username  string  `json:"username"`
	Password  string  `json:"password"`
}

type RegisterForm struct {
	Username  string   `json:"username"`
	Email	  string   `json:"email"`
	Password  string   `json:"password"`
}

type TokenResponse struct {
	AccessToken string  `json:"accesstoken"`
}

func main() {
	router := gin.Default()
	db := model.Connection()
	defer db.Close()
	model.Migrate(db)

	router.POST("/Login",LoginFunction)
	router.POST("/Register",RegisterFunction)
	router.Run(":8800")
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

func RegisterFunction(ctx *gin.Context) {
	
	var registerform RegisterForm
	ctx.BindJSON(&registerform)
	token := CreatToken()
	tokenrequest := TokenResponse{
		AccessToken: *token,
	}
	ctx.JSON(http.StatusOK,tokenrequest)
}

func LoginFunction(ctx *gin.Context) {
	
	var loginform LoginForm
	ctx.BindJSON(&loginform)
	if (loginform.Username == "hoge" && loginform.Password == "password") {
		token := CreatToken()
		tokenrequest := TokenResponse{
			AccessToken: *token,
		}
		ctx.JSON(http.StatusOK,tokenrequest)
	} else {
		error := Error {
			Error: "ログイン失敗",
		}
		ctx.JSON(http.StatusOK,error)
	}
}
