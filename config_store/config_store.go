package config_store

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type CommandStore = map[string]map[string]CommandInfo

type ConfigStore struct {
	Store *CommandStore
	SelectedWorkSpace string
	file *os.File
}

func InitStore() ConfigStore {
	data, file := readOrCreateFile()

	var commands *CommandStore

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
	for key := range *c.Store {
		fmt.Printf("- %s\n", key)
	}
}

func (c *ConfigStore) AddWorkspace(workspace string) {
	_, exists := (*c.Store)[workspace]

	if exists {
		log.Println("A workspace with this name already exists")
		return
	}

	(*c.Store)[workspace] = make(map[string]CommandInfo)
	// TODO: ensure that you can write to file
	updateConfigFile(c.file, c.Store)
}

func (c *ConfigStore) DeleteWorkspace(workspace string) {
	_, exists := (*c.Store)[workspace]

	if !exists {
		log.Println("A workspace with this name could not be found")
		return
	}

	delete(*c.Store, workspace)
	// TODO: ensure that you can write to file
	updateConfigFile(c.file, c.Store)
}

func (c *ConfigStore) SelectWorkspace(workspace string) {
	_, exists := (*c.Store)[workspace]

	if !exists {
		log.Println("A workspace with this name already exists")
		return
	}

	c.SelectedWorkSpace = workspace
}