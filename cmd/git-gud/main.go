package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid amount of arguments")
		return
	}
	fmt.Println("Intializing Banger")
}
