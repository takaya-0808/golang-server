package bookshelf

import(
	"bookself/app/util"
	"github.com/gin-gonic/gin"
	jwt "github.com/dgrijalva/jwt-go"
)

func GetArticl(c *gin.Context) {

	token, err := util.CheckToken(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "401"})
	} else {
		claims := token.Claims.(jwt.MapClaims)
		var msg = claims["user"]
		c.JSON(200, gin.H{"message": msg})
	}
}

func PutArticl() {}

func DeleteArticl() {}

func CreateAritcl() {}