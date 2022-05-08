package main

import (
	"io"
	"log"
	"os"
	"encoding/json"
	"github.com/durid-ah/go-flags"
)

type CommandStore = map[string]map[string]string

type Options struct {
	// Example of verbosity with level
	Verbose []bool `short:"v" long:"verbose" description:"Verbose output"`

	// Example of optional value
	User string `short:"u" long:"user" description:"User name" optional:"yes" optional-value:"pancake"`

	// Example of map with multiple default values
	Users map[string]string `long:"users" description:"User e-mail map" default:"system:system@example.org" default:"admin:admin@example.org"`
}

var options Options

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

func PopulateConfigFile(file *os.File) *CommandStore {
	commands := make(CommandStore)
	commands["main"] = make(map[string]string)
	jsonData, _ := json.Marshal(commands)
	file.Write(jsonData)
	return &commands
}

func main() {
	log.Println("Getting started")
	data, file := ReadOrCreateFile()
	var commands *CommandStore
	
	if len(data) == 0 {
		commands = PopulateConfigFile(file)
	} else {
		json.Unmarshal(data, &commands)
		log.Println(commands)
	}

	// var parser = flags.NewParser(&options, os.Args)
	flags.ParseArgs(&options, os.Args)
	log.Println(options)
}
