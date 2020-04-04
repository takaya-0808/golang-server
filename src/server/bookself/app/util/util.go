package util

import(
	"time"
	"bookself/app/login"
	jwt "github.com/dgrijalva/jwt-go"
    "github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

var secretKey = "75c92a074c341e9964329c0550c2673730ed8479c885c43122c90a2843177d5ef21cb50cfadcccb20aeb730487c11e09ee4dbbb02387242ef264e74cbee97213"

func CreateToken(c *gin.Context) {
	
	Flag := login.CheckLogin(c)
	
	if Flag != nil {
		token := jwt.New(jwt.GetSigningMethod("HS256"))
		token.Claims = jwt.MapClaims {
			"user": "hoge",
			"exp":  time.Now().Add(time.Hour * 1).Unix(),
		}
		tokenString, err := token.SignedString([]byte(secretKey))
		if err != nil {
			c.JSON(500, gin.H{"message": "Token don't Create"})
		} else {
			c.JSON(200, gin.H{"token": tokenString})
		}
	} else {
		c.JSON(500, gin.H{"message": "Could not generate token"})
	}
}


func CheckToken(c *gin.Context) (*jwt.Token, error) {
	
	token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		b := []byte(secretKey)
		return b, nil
	})
	return token, err
}