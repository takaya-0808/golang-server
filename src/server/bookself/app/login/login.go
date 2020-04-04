package login

import(
	"github.com/gin-gonic/gin"
)

type loginForm struct {
	UserName   string  `json:"username"`
	PassWord   string  `json:"password"`
}

func CheckLogin(c *gin.Context) *loginForm {

	var loginform loginForm
	c.BindJSON(&loginform)
	if loginform.UserName == "hoge" && loginform.PassWord == "password" {
		return &loginform
	} else {
		return nil
	}
}