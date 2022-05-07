package main

import (
	"io"
	"log"
	"os"
	"encoding/json"
)

func fatalLog(err error) {
	if err != nil {
		log.Fatal("Read file failed ", err)
	}
}

func ReadOrCreateFile() ([]byte, *os.File) {
	file, err := os.OpenFile("./config.json", os.O_CREATE | os.O_RDWR, 0666)
	fatalLog(err)

	data, err := io.ReadAll(file)
	fatalLog(err)

	return data, file
}

func main() {
	log.Println("Getting started")
	data, file := ReadOrCreateFile()
	commands := make(map[string]map[string]string)
	if len(data) == 0 {
		commands["main"] = make(map[string]string)
		jsonData, _ := json.Marshal(commands)
		file.Write(jsonData)
	} else {
		json.Unmarshal(data, &commands)
		log.Println(commands)
	}
}
