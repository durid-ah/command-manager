package main;

type CommandStore = map[string]map[string]string

type ConfigStore struct {
	store *CommandStore
}

func InitStore() ConfigStore {
	store := make(CommandStore)
	return ConfigStore {
		store: &store,
	}
}