package main

import (
	"fmt"
	"os"

	commands "github.com/F00bis/GitClone/cmd/main"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "The number of command line arguments is too small")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		commands.Init()
	case "catfile":
		commands.CatFile()
	default:
		fmt.Printf("Command Not Recognized")
	}

}
