package parser

import (
	"log"
	"strings"
	"github.com/durid-ah/command-manager/config_store"
)


func ParseCommands(
	input *string, commands *config_store.ConfigStore) {
	values := strings.Split(*input, " ")

	switch(values[0]) {
		case "--add":
			if len(values) >= 4 {
				commands.AddCommand(values[1], values[2], values[3:])
				return
			}
			log.Println("Adding command alias requires at least three argument")
	
		case "--del":
			if len(values) == 2 {
				commands.DeleteCommand(values[1])
				return
			}
			log.Println("Deleting command alias requires only one argument")

		case "--list":
			if len(values) == 1 {
				commands.ListCommandsForWorkspace("")
				return
			} else if len(values) == 2 {
				commands.ListCommandsForWorkspace(values[1])
				return
			}
			log.Println("Deleting command alias requires only one argument")

		case "--add-ws":
			if len(values) == 2 {
				commands.AddWorkspace(values[1])
				return
			}
			log.Println("Adding workspace requires one argument")
		
		case "--sel-ws":
			if len(values) == 2 {
				commands.SelectWorkspace(values[1])
				return
			}
			log.Println("Adding workspace requires one argument")

		case "--del-ws":
			if len(values) != 2 {
				log.Println("Deleting workspace requires one argument")
			} else if values[1] == "main" {
				log.Println("Deleting main workspace is not allowed")
			} else {
				commands.DeleteWorkspace(values[1])
			}

		case "--list-ws":
			if len(values) == 1 {
				commands.ListWorkspaces()
				return
			}
			log.Println("Listing workspace does not take any arguments")
		
		case "--help":
			ShowAppCommands()

		default:
			commands.ExecCommand(values[0])
	}
}
