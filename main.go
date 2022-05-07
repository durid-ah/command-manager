package main

import (
	"io"
	"log"
	"os"
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
	data, _ := ReadOrCreateFile()
	if len(data) == 0 {
		log.Println("Instantiate a map and write to file")
	} else {
		log.Println("Parse the data to a map")
	}
}
