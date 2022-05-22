package parser

import (
	"strings"
	"github.com/durid-ah/command-manager/config_store"
)


func ParseCommands(
	input *string, commands *config_store.ConfigStore) {
	println("ParseCommands")
	println(*input)
	values := strings.Split(*input, " ")

	switch(values[0]) {
		case "--add-ws":
			println(values[1])
			// TODO: check that there is another command
			commands.AddWorkspace(
				strings.TrimRight(values[1], "\r\n"))
	}
}