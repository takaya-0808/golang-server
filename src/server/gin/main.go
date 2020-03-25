package main

import(
	"gin/config/setup"
	"gin/config/controllers"
)

func main() {
	db := setup.Connection()
	controllers.Router(db)
	defer db.Close()	
}


