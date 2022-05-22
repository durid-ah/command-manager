package config_store

import(
	"log"
	"os"
	"io"
	"encoding/json"
)

func fatalLog(err error) {
	if err != nil {
		log.Fatal("Read file failed ", err)
	}
}

func readOrCreateFile() ([]byte, *os.File) {
	file, err := os.OpenFile("./config.json", os.O_CREATE | os.O_RDWR, 0666)
	fatalLog(err)

	data, err := io.ReadAll(file)
	fatalLog(err)

	return data, file
}

func populateConfigFile(file *os.File) *CommandStore {
	commands := make(CommandStore)
	commands["main"] = make(map[string]string)
	jsonData, _ := json.Marshal(commands)
	file.Write(jsonData)
	return &commands
}