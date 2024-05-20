package commands

import (
	"fmt"
	"os"
)

func Init() {
	fmt.Println("Starting Git initialization")

	for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating file: %s\n", dir)
		}
	}

	headFileContent := []byte("ref: refs/heads/main\n")

	if err := os.WriteFile(".git/HEAD", headFileContent, 0644); err != nil {
		fmt.Printf("Error writing content to .git/HEAD\n")
	}

	fmt.Printf("Initialized Git Repository\n")
}
