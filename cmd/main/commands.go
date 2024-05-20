package commands

import (
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"strings"
)

func Init() {
	fmt.Println("Starting Git initialization")

	for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating file: %s\n", dir)
			os.Exit(1)
		}
	}

	headFileContent := []byte("ref: refs/heads/main\n")

	if err := os.WriteFile(".git/HEAD", headFileContent, 0644); err != nil {
		fmt.Printf("Error writing content to .git/HEAD\n")
		os.Exit(1)
	}

	fmt.Printf("Initialized Git Repository\n")
}

func CatFile() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Expected object hash")
		os.Exit(1)
	}

	sha := os.Args[2]

	path := fmt.Sprintf(".git/objects/%s/%s", sha[0:2], sha[2:])

	file, err := os.Open(path)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %s: %s", path, err)
		os.Exit(1)
	}

	reader, _ := zlib.NewReader(io.Reader(file))

	content, err := io.ReadAll(reader)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %s", path, err)
		os.Exit(1)
	}

	parts := strings.Split(string(content), "\x00")

	fmt.Print(parts[1])

	reader.Close()
}
