package config_store

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"unicode/utf8"
)

type CommandStore = map[string]map[string]CommandInfo

type ConfigStore struct {
	Store CommandStore
	SelectedWorkSpace string
	file *os.File
}

func InitStore() ConfigStore {
	data, file := readOrCreateFile()

	var commands CommandStore

	println("InitStore")
	fmt.Printf("Stuff: %s \n", data)

	if len(data) == 0 {
		commands = populateConfigFile(file)
	} else {
		json.Unmarshal(data, &commands)
	}

	return ConfigStore{
		Store: commands, SelectedWorkSpace: "main", file: file,
	}
}

func (c *ConfigStore) ListWorkspaces() {
	fmt.Println("Workspaces:")
	for key := range c.Store {
		fmt.Printf("- %s\n", key)
	}
}

func (c *ConfigStore) AddWorkspace(workspace string) {
	_, exists := c.Store[workspace]

	if exists {
		log.Println("A workspace with this name already exists")
		return
	}

	c.Store[workspace] = make(map[string]CommandInfo)
	// TODO: ensure that you can write to file
	updateConfigFile(c.file, c.Store)
}

func (c *ConfigStore) DeleteWorkspace(workspace string) {
	_, exists := c.Store[workspace]

	if !exists {
		log.Println("A workspace with this name could not be found")
		return
	}

	delete(c.Store, workspace)
	// TODO: ensure that you can write to file
	updateConfigFile(c.file, c.Store)
}

func (c *ConfigStore) SelectWorkspace(workspace string) {
	_, exists := c.Store[workspace]

	if !exists {
		log.Println("A workspace with this name doesn't exists")
		return
	}

	c.SelectedWorkSpace = workspace
}

func (c *ConfigStore) AddCommand(alias string, cwd string, command []string) {
	_, exists := c.Store[c.SelectedWorkSpace][alias]

	if exists {
		log.Printf("An alias {%s} under workspace {%s} already exists", alias, c.SelectedWorkSpace)
		return
	}

	
	c.Store[c.SelectedWorkSpace][alias] = CommandInfo{Command: strings.Join(command, " "), Cwd: cwd}
	updateConfigFile(c.file, c.Store)
}

func (c *ConfigStore) DeleteCommand(alias string) {
	_, exists := c.Store[c.SelectedWorkSpace][alias]

	if !exists {
		log.Printf("An alias {%s} under workspace {%s} doesn't exists", alias, c.SelectedWorkSpace)
		return
	}

	delete(c.Store[c.SelectedWorkSpace], alias)
	updateConfigFile(c.file, c.Store)
}

func (c *ConfigStore) ListCommandsForWorkspace(workspace string) {
	var workspaceName string;
	if workspace == "" {
		workspaceName = c.SelectedWorkSpace
	} else {
		workspaceName = workspace
	}

	workspaceMap := c.Store;
	for key, val := range workspaceMap[workspaceName] {
		fmt.Printf("{%s} => {%s} \n", key, val.Command)
	}
}

func (c *ConfigStore) ExecCommand(alias string) {
	_, exists := c.Store[c.SelectedWorkSpace][alias]

	if !exists {
		log.Printf("An alias {%s} under workspace {%s} doesn't exists", alias, c.SelectedWorkSpace)
		return
	}

	commandInfo := c.Store[c.SelectedWorkSpace][alias]
	cmd := createCommand(&commandInfo.Command)
	
	// Setup the command
	cmd.Dir = commandInfo.Cwd	
	outPipe, _ :=  cmd.StdoutPipe()
	errPipe, _ :=  cmd.StderrPipe()
	out := io.MultiReader(outPipe, errPipe)
	
	startErr := cmd.Start()
	if startErr != nil {
		log.Printf("%s", startErr)
	}

	// Logic to handle command cancellation
	cancelSignal := make(chan os.Signal, 1)
	signal.Notify(cancelSignal, os.Interrupt)
	defer signal.Reset(os.Interrupt)
	defer close(cancelSignal)
	
	// Stream the command execution output
	outputBuffer := make([]byte, utf8.UTFMax)
	
	readLoop: for {
		select {
			case <- cancelSignal:
				break readLoop
			default:
				n, err := out.Read(outputBuffer)
				fmt.Printf("%s", outputBuffer[:n])
				if err == io.EOF {
					break readLoop
				}
		}

	}

	cmd.Wait()
}