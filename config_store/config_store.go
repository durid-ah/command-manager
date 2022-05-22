package config_store

import "encoding/json"

type CommandStore = map[string]map[string]string

type ConfigStore struct {
	Store *CommandStore
}

func InitStore() ConfigStore {
	data, file := readOrCreateFile()

	var commands *CommandStore

	if len(data) == 0 {
		commands = populateConfigFile(file)
	} else {
		json.Unmarshal(data, &commands)
	}

	return ConfigStore{
		Store: commands,
	}
}