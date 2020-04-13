package main

import(
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"Register/User/controller"
	"Register/User/model"
)

func main()  {
	
	db := initDB()
	defer db.Close()
	db.LogMode(true)

	router := gin.Default()
	controller.Router(router)
}

func initDB() *gorm.DB {
	db,err := gorm.Open("mysql","root:password@tcp(godockerDB)/test2?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	db.AutoMigrate(&model.UserData{})
	return db
}
