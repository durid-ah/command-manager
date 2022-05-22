package main

import (
	"bufio"
	"log"
	"os"
	"github.com/durid-ah/command-manager/config_store"
	"github.com/durid-ah/command-manager/parser"
)

// func fatalLog(err error) {
// 	if err != nil {
// 		log.Fatal("Read file failed ", err)
// 	}
// }

func main() {
	log.Println("Getting started")
	configs := config_store.InitStore()

	// for key, _ := range *configs.Store {
	// 	println(key)
	// }

	for {
		inputReader := bufio.NewReader(os.Stdin)
		input, _ :=  inputReader.ReadString('\n')
		parser.ParseCommands(&input, &configs)

		for key, _ := range *configs.Store {
			println(key)
		}

		break
	}
}
