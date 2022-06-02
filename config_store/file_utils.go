package config_store

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func fatalLog(err error) {
	if err != nil {
		log.Fatal("Read file failed ", err)
	}
}

func readOrCreateFile() ([]byte, *os.File) {
	file, err := os.OpenFile("./config.json", os.O_CREATE | os.O_RDWR, 0755)
	fatalLog(err)

	data, err := io.ReadAll(file)
	fatalLog(err)

	return data, file
}

func populateConfigFile(file *os.File) CommandStore {
	commands := make(CommandStore)
	commands["main"] = make(map[string]CommandInfo)
	updateConfigFile(file, commands)
	return commands
}

func updateConfigFile(file *os.File, commands CommandStore) {
	jsonData, _ := json.Marshal(commands)
	file.Truncate(0)
	file.Seek(0,0)
	file.Write(jsonData)
}