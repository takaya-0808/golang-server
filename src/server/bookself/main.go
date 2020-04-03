package main

import(
	
	"bookself/controller"
)

func main() {
	r := controller.Router()
	controller.Routers(r)
}