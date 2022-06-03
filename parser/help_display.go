package parser

import "fmt"

func ShowAppCommands() {
	fmt.Println("")
	fmt.Println("(<workspace-name>)> The name between parentheses refers to the current workspace")
	fmt.Println("")

	fmt.Println("Workspace Actions:")
	fmt.Println("")

	fmt.Print("\t --add-ws <worskspace-name>:")
	fmt.Print(" Create a new workspace with the selected \n")

	fmt.Print("\t --sel-ws <worskspace-name>:")
	fmt.Print(" Selects a workspace as the current working one \n")

	fmt.Print("\t --del-ws <worskspace-name>:")
	fmt.Print(" Deletes the specified workspace as long as it is not 'main' \n")

	fmt.Println("")
	fmt.Println("Command Actions:")
	fmt.Println("")

	fmt.Print("\t --add <alias> <cwd> <command[]>:")
	// TODO: Finish redoing
	fmt.Print(" Create a new workspace with the selected \n")
}