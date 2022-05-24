package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"github.com/durid-ah/command-manager/config_store"
	"github.com/durid-ah/command-manager/parser"
)

func main() {
	log.Println("Getting started")
	configs := config_store.InitStore()

	for {
		inputReader := bufio.NewReader(os.Stdin)
		input, _ :=  inputReader.ReadString('\n')
		commandString := strings.TrimRight(input, "\r\n")
		if commandString == "--quit" {
			break
		}
		parser.ParseCommands(&commandString, &configs)

	}
}
