package commands

import (
	"fmt"
	"git-gud/internal/repository"
)

var commandMap = map[string]func(){
	"init":   Init,
	"commit": Commit,
	"status": Status,
	"add":    Add,
}

func ExecuteCommand(cmd []string) {
	if action, found := commandMap[cmd[0]]; found {
		action()
	} else {
		fmt.Println("Unknown command:", cmd[0])
	}
}

func Init() {

	directories := []string{
		".gg/objects",
		".gg/refs/heads",
		".gg/refs/tags",
	}
	for _, directory := range directories {
		repository.CreateDirectory(directory)
	}
	repository.WriteFile(".gg/HEAD", "ref: refs/heads/main")

}

func Add() {

}

func Commit() { fmt.Println("Committing...") }
func Status() { fmt.Println("Status...") }
