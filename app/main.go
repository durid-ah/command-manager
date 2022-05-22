package main

import (
	"log"
	"github.com/durid-ah/command-manager/config_store"
)


// func fatalLog(err error) {
// 	if err != nil {
// 		log.Fatal("Read file failed ", err)
// 	}
// }

func main() {
	log.Println("Getting started")
	configs := config_store.InitStore()
	main := (*configs.Store)["main"]
	stuff := main["stuff"]
	println(stuff)
	println((*configs.Store))
}
