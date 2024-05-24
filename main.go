package main

import (
	"golang-template/router"
	"log"
)

func main() {
	router := router.NewRouter()
	log.Fatalln(router.Start(":" + "8000"))
}
