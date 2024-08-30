package commands

import (
	"fmt"
	"git-gud/internal/repository"
)

var commandMap = map[string]func([]string){
	"init":   Init,
	"commit": Commit,
	"status": Status,
	"add":    Add,
}

func ExecuteCommand(cmd []string) {
	if action, found := commandMap[cmd[0]]; found {
		action(cmd[1:])
	} else {
		fmt.Println("Unknown command:", cmd[0])
	}
}

func Init(args []string) {
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

func Commit(args []string) {
	fmt.Println("Committing...")
}

func Status(args []string) {
	fmt.Println("Status...")
}

func Add(args []string) {

	if len(args) == 0 {
		fmt.Println("No files specified to add.")
		return
	}
	if args[0] == "." || args[0] == "-all" || args[0] == "-A" {
		addAll()
	} else {
		addSpecified(args)
	}
}

// HELPERS

func addAll() {
	files := repository.ReadDirectoryFiles()
	for _, file := range files {
		err := handleFileAdd(file.Path)
		if err != nil {
			fmt.Println("Error adding file:", err)
		}
	}
}

func addSpecified(files []string) {
	for _, file := range files {
		err := handleFileAdd(file)
		if err != nil {
			fmt.Println("Error adding file:", err)
		}
	}
}
