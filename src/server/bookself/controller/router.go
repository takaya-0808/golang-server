package controller

import(
	"github.com/gin-gonic/gin"
	"bookself/app/util"
	"bookself/app/bookshelf"
)

func Router() *gin.Engine {
	router := gin.Default()
	return router
}

func Routers(router *gin.Engine) {

	ArticlService := router.Group("/api/v1/articl")
	{
		ArticlService.GET("/:id", bookshelf.GetArticl)
		//ArticlService.POST("", bookshelf.CreateAritcl)
		//ArticlService.PUT("/:id", bookshelf.PutArticl)
		//ArticlService.DELETE("/:id", bookshelf.DeleteArticl)
	}

 	LoginService := router.Group("/api/v1/login")
	{
		LoginService.POST("", util.CreateToken)
	}

/*	RegisterService := router.Group("/api/v1/register")
	{
		RegisterService.POST("", util.CreateToken)
	}
*/
	router.Run(":8080")
}


