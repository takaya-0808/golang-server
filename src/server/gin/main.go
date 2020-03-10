package main

import(
	"net/http"	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gin/model"
	"github.com/jinzhu/gorm"
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


type Handle struct {
	ID        int      `json:"id"`
	Username  string   `json:"username"`
	Email	  string   `json:"email"`
	Password  string   `json:"password"`
}

func New(db *gorm.DB) *Handle {
	return  &Handle{}
}

func main() {
	router := gin.Default()
	db := model.Connection()
	defer db.Close()
	model.Migrate(db)
	hdl := New(db)

	router.POST("/Login",hdl.LoginFunction)
	router.POST("/Register",hdl.RegisterFunction)
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

func (h *Handle)RegisterFunction(ctx *gin.Context) {
	
	var registerform RegisterForm
	ctx.BindJSON(&registerform)
	model.NewInsert(registerform.Username, registerform.Email, registerform.Password, h.db)
	token := CreatToken()
	tokenrequest := TokenResponse{
		AccessToken: *token,
	}
	ctx.JSON(http.StatusOK,tokenrequest)
}

func (h *Handle)LoginFunction(ctx *gin.Context) {
	
	var loginform LoginForm
	ctx.BindJSON(&loginform)
	userID := model.LoginProcess(loginform.Username, loginform.Password, h.db)
	if (userID != nil) {
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

